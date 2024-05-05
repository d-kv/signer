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
	if err != nil {
		return err
	}
	return err
}
