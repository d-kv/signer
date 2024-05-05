package usecase

import (
	"context"
	"d-kv/signer/db-common/entity"
)

type (
	TenantRepo interface {
		Create(context.Context, *entity.Tenant) error
		FindById(context.Context, string) (entity.Tenant, error)
		Update(context.Context, *entity.Tenant) error
		DeleteById(context.Context, string) error
		FindAll(context.Context) ([]entity.Tenant, error)
	}

	DeviceRepo interface {
		Create(context.Context, *entity.Device) error
		FindById(context.Context, string) (entity.Device, error)
		Update(context.Context, *entity.Device) error
		DeleteById(context.Context, string) error
		FindAll(context.Context) ([]entity.Device, error)
	}

	IntegrationRepo interface {
		Create(context.Context, *entity.Integration) error
		FindById(context.Context, string) (entity.Integration, error)
		Update(context.Context, *entity.Integration) error
		DeleteById(context.Context, string) error
		FindAll(context.Context) ([]entity.Integration, error)
	}

	ProfileRepo interface {
		Create(context.Context, *entity.Profile) error
		FindById(context.Context, string) (entity.Profile, error)
		Update(context.Context, *entity.Profile) error
		DeleteById(context.Context, string) error
		FindAll(context.Context) ([]entity.Profile, error)
	}

	BundleIdRepo interface {
		Create(context.Context, *entity.BundleId) error
		FindById(context.Context, string) (entity.BundleId, error)
		FindByIntegrationId(context.Context, string) (entity.BundleId, error)
		Update(context.Context, *entity.BundleId) error
		DeleteById(context.Context, string) error
		FindAll(context.Context) ([]entity.BundleId, error)
	}

	CertificateRepo interface {
		Create(context.Context, *entity.Certificate) error
		FindById(context.Context, string) (entity.Certificate, error)
		Update(context.Context, *entity.Certificate) error
		DeleteById(context.Context, string) error
		FindAll(context.Context) ([]entity.Certificate, error)
	}

	UserRepo interface {
		Create(context.Context, *entity.User) error
		FindById(context.Context, string) (entity.User, error)
		Update(context.Context, *entity.User) error
		DeleteById(context.Context, string) error
		FindAll(context.Context) ([]entity.User, error)
	}

	CapabilityRepo interface {
		Create(context.Context, *entity.Capability) error
		FindByBundleIdId(context.Context, string) ([]entity.Capability, error)
		Update(context.Context, *entity.Capability) error
		DeleteById(context.Context, string) error
		FindAll(context.Context) ([]entity.Capability, error)
	}
)

type CommandRepo interface {
	CreateBundleIdCommand(context.Context, entity.CreateBundleId) (error, entity.CreateBundleId)
	CreateDeviceCommand(context.Context, entity.CreateDevice) (error, entity.CreateDevice)
	CreateEnableCapabilityTypeCommand(context.Context, entity.EnableCapabilityType) (error, entity.EnableCapabilityType)

	FindByStatusBundleIdCommand(context.Context, entity.Status) []entity.CreateBundleId
	FindByStatusDeviceCommand(context.Context, entity.Status) []entity.CreateDevice
	FindByStatusEnableCapabilityTypeCommand(context.Context, entity.Status) []entity.EnableCapabilityType

	GetStatusByIdBundleIdCommand(context.Context, uint) (error, entity.Status)
	GetStatusByIdDeviceCommand(context.Context, uint) (error, entity.Status)
	GetStatusByIdEnableCapabilityTypeCommand(context.Context, uint) (error, entity.Status)

	SetStatusByIdBundleIdCommand(context.Context, uint, entity.Status) error
	SetStatusByIdDeviceCommand(context.Context, uint, entity.Status) error
	SetStatusByIdEnableCapabilityTypeCommand(context.Context, uint, entity.Status) error
}

type VaultRepo interface {
	FindTokenByIntegrationId(context.Context, string) (error, *entity.IntegrationToken)
	SaveIntegrationToken(context.Context, *entity.IntegrationToken) error
}
