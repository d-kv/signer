package usecase

import (
	"command-executor/internal/entity"
	"context"
)

type (
	TenantRepo interface {
		Create(context.Context, *entity.Tenant) error
		FindById(context.Context, uint) (entity.Tenant, error)
		Update(context.Context, *entity.Tenant) error
		DeleteById(context.Context, uint) error
	}

	DeviceRepo interface {
		Create(context.Context, *entity.Device) error
		FindById(context.Context, uint) (entity.Device, error)
		Update(context.Context, *entity.Device) error
		DeleteById(context.Context, uint) error
	}

	IntegrationRepo interface {
		Create(context.Context, *entity.Integration) error
		FindById(context.Context, uint) (entity.Integration, error)
		Update(context.Context, *entity.Integration) error
		DeleteById(context.Context, uint) error
	}

	ProfileRepo interface {
		Create(context.Context, *entity.Profile) error
		FindById(context.Context, uint) (entity.Profile, error)
		Update(context.Context, *entity.Profile) error
		DeleteById(context.Context, uint) error
	}

	BundleIdRepo interface {
		Create(context.Context, *entity.BundleId) error
		FindById(context.Context, uint) (entity.BundleId, error)
		Update(context.Context, *entity.BundleId) error
		DeleteById(context.Context, uint) error
	}

	CertificateRepo interface {
		Create(context.Context, *entity.Certificate) error
		FindById(context.Context, uint) (entity.Certificate, error)
		Update(context.Context, *entity.Certificate) error
		DeleteById(context.Context, uint) error
	}

	UserRepo interface {
		Create(context.Context, *entity.User) error
		FindById(context.Context, uint) (entity.User, error)
		Update(context.Context, *entity.User) error
		DeleteById(context.Context, uint) error
	}

	CapabilityRepo interface {
		Create(context.Context, *entity.Capability) error
		FindById(context.Context, uint) (entity.Capability, error)
		Update(context.Context, *entity.Capability) error
		DeleteById(context.Context, uint) error
	}
)

type CommandRepo interface {
	FindByStatusBundleIdCommand(context.Context, entity.Status) []entity.CreateBundleId
	FindByStatusDeviceCommand(context.Context, entity.Status) []entity.CreateDevice
	FindByStatusEnableCapabilityTypeCommand(context.Context, entity.Status) []entity.EnableCapabilityType

	SetStatusByIdBundleIdCommand(context.Context, uint, entity.Status) error
	SetStatusByIdDeviceCommand(context.Context, uint, entity.Status) error
	SetStatusByIdEnableCapabilityTypeCommand(context.Context, uint, entity.Status) error
}
