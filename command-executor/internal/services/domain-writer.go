package services

import (
	"context"
	"d-kv/signer/command-executor/pkg/entity"
	dbEntity "d-kv/signer/db-common/entity"
	"d-kv/signer/db-common/repo/domain"
	"errors"
	"gorm.io/gorm"
)

type DataBaseService struct {
	Repo *domain.PostgresDomainRepo
}

func NewDataBaseService(repo *domain.PostgresDomainRepo) *DataBaseService {
	return &DataBaseService{Repo: repo}
}

func (s *DataBaseService) WriteCapability(ctx context.Context, err error, operation dbEntity.EnableCapabilityType) error {
	id, err := s.Repo.BundleIdRepo.FindById(ctx, operation.BundleId)
	if err != nil {
		return err
	}
	converted := entity.ConvertCapability(&id, &operation)
	err = s.Repo.CapabilityRepo.Create(ctx, converted)
	if err != nil {
		return err
	}

	return err
}
func (s *DataBaseService) WriteDevice(ctx context.Context, operation dbEntity.CreateDevice) error {
	userid, err := s.Repo.UserRepo.FindById(ctx, operation.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			userid = dbEntity.User{
				ID:   operation.UserID,
				Name: operation.UserName,
			}
		} else {
			return err
		}
	}
	deviceId, err := s.Repo.DeviceRepo.FindById(ctx, operation.DeviceUdid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			var profiles []dbEntity.Profile
			var integrations []dbEntity.Integration
			integration, err1 := s.Repo.IntegrationRepo.FindById(ctx, operation.IntegrationId)
			if err1 != nil {
				return err1
			}
			integrations = append(integrations, integration)
			converted := entity.ConvertDevice(&operation, &userid, profiles, integrations)
			err = s.Repo.DeviceRepo.Create(ctx, converted)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		integration, err2 := s.Repo.IntegrationRepo.FindById(ctx, operation.IntegrationId)
		if err2 != nil {
			return err2
		}
		deviceId.Integrations = append(deviceId.Integrations, integration)
		err = s.Repo.DeviceRepo.Update(ctx, &deviceId)
		if err != nil {
			return err
		}
	}
	return err
}

func (s *DataBaseService) WriteBundleId(ctx context.Context, operation dbEntity.CreateBundleId, response *entity.BundleIdResponse) error {
	id, err := s.Repo.IntegrationRepo.FindById(ctx, operation.IntegrationId)
	if err != nil {
		return err
	}
	converted := entity.ConvertBundleId(&id, &operation, response)
	err = s.Repo.BundleIdRepo.Create(ctx, converted)
	return err
}

func (s *DataBaseService) WriteProfile(ctx context.Context, operation dbEntity.CreateProfile, resp *entity.ProfileResponse) error {
	var devices []dbEntity.Device
	var certificates []dbEntity.Certificate
	for _, deviceId := range operation.DeviceIds {
		device, err := s.Repo.DeviceRepo.FindById(ctx, deviceId)
		if err != nil {
			return err
		}
		devices = append(devices, device)
	}
	for _, id := range operation.CertificateIds {
		certificate, err := s.Repo.CertificateRepo.FindById(ctx, id)
		if err != nil {
			return err
		}
		certificates = append(certificates, certificate)
	}
	bundleId, err := s.Repo.BundleIdRepo.FindById(ctx, operation.BundleId)
	if err != nil {
		return err
	}
	integration, err := s.Repo.IntegrationRepo.FindById(ctx, operation.IntegrationId)
	if err != nil {
		return err
	}
	converted := entity.ConvertProfile(&operation, devices, certificates, resp, &bundleId, &integration)
	err = s.Repo.ProfileRepo.Create(ctx, converted)
	if err != nil {
		return err
	}
	err = s.UpdateArrayDevices(ctx, devices, converted, err)
	if err != nil {
		return err
	}
	err = s.UpdateArrayCertificates(ctx, certificates, converted, err)
	if err != nil {
		return err
	}
	return s.Repo.BundleIdRepo.Update(ctx, &bundleId)
}

func (s *DataBaseService) WriteCertificate(ctx context.Context, operation dbEntity.CreateCertificate, resp *entity.CertificateResponse) error {
	var profiles []dbEntity.Profile
	integration, err := s.Repo.IntegrationRepo.FindById(ctx, operation.IntegrationId)
	if err != nil {
		return err
	}
	converted := entity.ConvertCertificate(&operation, profiles, &integration, resp)
	err = s.Repo.CertificateRepo.Create(ctx, converted)
	return err
}

func (s *DataBaseService) UpdateArrayDevices(ctx context.Context, devices []dbEntity.Device, converted *dbEntity.Profile, err error) error {
	for _, device := range devices {
		device.Profiles = append(device.Profiles, *converted)
		err = s.Repo.DeviceRepo.Update(ctx, &device)
		if err != nil {
			return err
		}
	}
	return err
}

func (s *DataBaseService) UpdateArrayCertificates(ctx context.Context, certificates []dbEntity.Certificate, converted *dbEntity.Profile, err error) error {
	for _, cert := range certificates {
		cert.Profiles = append(cert.Profiles, *converted)
		err = s.Repo.CertificateRepo.Update(ctx, &cert)
		if err != nil {
			return err
		}
	}
	return err
}
