package repo

import (
	"context"
	pkgEntity "d-kv/signer/command-executor/pkg/entity"
	"d-kv/signer/db-common/entity"
	"d-kv/signer/db-common/repo/command"
	"d-kv/signer/db-common/repo/domain"
	"errors"
	"gorm.io/gorm"
)

func WriteCapability(ctx context.Context, queue *command.Repo, repo *domain.PostgresDomainRepo, err error, operation entity.EnableCapabilityType) error {
	id, err := repo.BundleIdRepo.FindById(ctx, operation.BundleId)
	if err != nil {
		return err
	}
	converted := pkgEntity.ConvertCapability(&id, &operation)
	err = repo.CapabilityRepo.Create(ctx, converted)
	if err != nil {
		return err
	}
	err = queue.SetStatusByIdEnableCapabilityTypeCommand(ctx, operation.ID, entity.Completed)
	return err
}
func WriteDevice(ctx context.Context, queue *command.Repo, repo *domain.PostgresDomainRepo, operation entity.CreateDevice) error {
	userid, err := repo.UserRepo.FindById(ctx, operation.UserID)
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
	deviceId, err := repo.DeviceRepo.FindById(ctx, operation.DeviceUdid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			var profiles []entity.Profile
			var integrations []entity.Integration
			integration, err1 := repo.IntegrationRepo.FindById(ctx, operation.IntegrationId)
			if err1 != nil {
				return err1
			}
			integrations = append(integrations, integration)
			converted := pkgEntity.ConvertDevice(&operation, &userid, profiles, integrations)
			err = repo.DeviceRepo.Create(ctx, converted)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		integration, err2 := repo.IntegrationRepo.FindById(ctx, operation.IntegrationId)
		if err2 != nil {
			return err2
		}
		deviceId.Integrations = append(deviceId.Integrations, integration)
		err = repo.DeviceRepo.Update(ctx, &deviceId)
		if err != nil {
			return err
		}
		err = queue.SetStatusByIdDeviceCommand(ctx, operation.ID, entity.Completed)
		if err != nil {
			return err
		}

	}
	return err
}

func WriteBundleId(ctx context.Context, queue *command.Repo, repo *domain.PostgresDomainRepo, operation entity.CreateBundleId) error {
	id, err := repo.IntegrationRepo.FindById(ctx, operation.IntegrationId)
	if err != nil {
		return err
	}
	converted := pkgEntity.ConvertBundleId(&id, &operation)
	err = repo.BundleIdRepo.Create(ctx, converted)
	if err != nil {
		return err
	}
	err = queue.SetStatusByIdBundleIdCommand(ctx, operation.ID, entity.Completed)
	if err != nil {
		return err
	}
	return err
}
