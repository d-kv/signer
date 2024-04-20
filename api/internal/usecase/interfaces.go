package usecase

import (
	"context"
	"d-kv/signer/api/internal/entity"
)

type CommandRepo interface {
	CreateBundleIdCommand(context.Context, *entity.CreateBundleId) error
	CreateDeviceCommand(context.Context, *entity.CreateDevice) error
	CreateEnableCapabilityTypeCommand(context.Context, *entity.EnableCapabilityType) error
}
