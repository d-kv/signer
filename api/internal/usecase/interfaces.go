package usecase

import (
	"api/internal/entity"
	"context"
)

type CommandRepo interface {
	CreateBundleIdCommand(context.Context, *entity.CreateBundleId) error
	CreateDeviceCommand(context.Context, *entity.CreateDevice) error
	CreateEnableCapabilityTypeCommand(context.Context, *entity.EnableCapabilityType) error
}
