package usecase

import (
	"command-executor/internal/entity"
	"context"
)

type CommandRepo interface {
	FindByStatusBundleIdCommand(context.Context, entity.Status) []entity.CreateBundleId
	FindByStatusDeviceCommand(context.Context, entity.Status) []entity.CreateDevice
	FindByStatusEnableCapabilityTypeCommand(context.Context, entity.Status) []entity.EnableCapabilityType

	SetStatusByIdBundleIdCommand(context.Context, uint, entity.Status) error
	SetStatusByIdDeviceCommand(context.Context, uint, entity.Status) error
	SetStatusByIdEnableCapabilityTypeCommand(context.Context, uint, entity.Status) error
}
