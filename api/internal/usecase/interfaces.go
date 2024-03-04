package usecase

import (
	"api/internal/entity"
	"context"
)

type CommandRepo interface {
	CreateBundleIdCommand(context.Context, entity.CreateBundleId) (entity.CreateBundleId, error)
	CreateDeviceCommand(context.Context, entity.CreateDevice) (entity.CreateDevice, error)
	CreateEnableCapabilityTypeCommand(context.Context, entity.EnableCapabilityType) (entity.EnableCapabilityType, error)
}
