package services

import (
	"context"
	pkgEntity "d-kv/signer/command-executor/pkg/entity"
	"d-kv/signer/db-common/entity"
	"errors"
	"gorm.io/gorm"
)

func (s *ProcessorService) WriteCapability(ctx context.Context, err error, operation entity.EnableCapabilityType) error {
	id, err := s.Repo.BundleIdRepo.FindById(ctx, operation.BundleId)
	if err != nil {
		return err
	}
	converted := pkgEntity.ConvertCapability(&id, &operation)
	err = s.Repo.CapabilityRepo.Create(ctx, converted)
	if err != nil {
		return err
	}
	err = s.Queue.SetStatusByIdEnableCapabilityTypeCommand(ctx, operation.ID, entity.Completed)
	return err
}
func (s *ProcessorService) WriteDevice(ctx context.Context, operation entity.CreateDevice) error {
	userid, err := s.Repo.UserRepo.FindById(ctx, operation.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			userid = entity.User{
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
			var profiles []entity.Profile
			var integrations []entity.Integration
			integration, err1 := s.Repo.IntegrationRepo.FindById(ctx, operation.IntegrationId)
			if err1 != nil {
				return err1
			}
			integrations = append(integrations, integration)
			converted := pkgEntity.ConvertDevice(&operation, &userid, profiles, integrations)
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
		err = s.Queue.SetStatusByIdDeviceCommand(ctx, operation.ID, entity.Completed)
		if err != nil {
			return err
		}

	}
	return err
}

func (s *ProcessorService) WriteBundleId(ctx context.Context, operation entity.CreateBundleId, response *pkgEntity.BundleIdResponse) error {
	id, err := s.Repo.IntegrationRepo.FindById(ctx, operation.IntegrationId)
	if err != nil {
		return err
	}
	converted := pkgEntity.ConvertBundleId(&id, &operation, response)
	err = s.Repo.BundleIdRepo.Create(ctx, converted)
	if err != nil {
		return err
	}
	err = s.Queue.SetStatusByIdBundleIdCommand(ctx, operation.ID, entity.Completed)
	if err != nil {
		return err
	}
	return err
}
