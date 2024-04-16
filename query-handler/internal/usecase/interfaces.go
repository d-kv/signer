package usecase

import (
	"context"
	"query-handler/internal/entity"
)

type (
	TenantRepo interface {
		Create(context.Context, *entity.Tenant) error
		FindById(context.Context, uint) (entity.Tenant, error)
		Update(context.Context, *entity.Tenant) error
		DeleteById(context.Context, uint) error
		FindAll(context.Context) ([]entity.Tenant, error)
	}

	DeviceRepo interface {
		Create(context.Context, *entity.Device) error
		FindById(context.Context, uint) (entity.Device, error)
		Update(context.Context, *entity.Device) error
		DeleteById(context.Context, uint) error
		FindAll(context.Context) ([]entity.Device, error)
	}

	IntegrationRepo interface {
		Create(context.Context, *entity.Integration) error
		FindById(context.Context, uint) (entity.Integration, error)
		Update(context.Context, *entity.Integration) error
		DeleteById(context.Context, uint) error
		FindAll(context.Context) ([]entity.Integration, error)
	}

	ProfileRepo interface {
		Create(context.Context, *entity.Profile) error
		FindById(context.Context, uint) (entity.Profile, error)
		Update(context.Context, *entity.Profile) error
		DeleteById(context.Context, uint) error
		FindAll(context.Context) ([]entity.Profile, error)
	}

	BundleIdRepo interface {
		Create(context.Context, *entity.BundleId) error
		FindById(context.Context, uint) (entity.BundleId, error)
		Update(context.Context, *entity.BundleId) error
		DeleteById(context.Context, uint) error
		FindAll(context.Context) ([]entity.BundleId, error)
	}

	CertificateRepo interface {
		Create(context.Context, *entity.Certificate) error
		FindById(context.Context, uint) (entity.Certificate, error)
		Update(context.Context, *entity.Certificate) error
		DeleteById(context.Context, uint) error
		FindAll(context.Context) ([]entity.Certificate, error)
	}

	UserRepo interface {
		Create(context.Context, *entity.User) error
		FindById(context.Context, uint) (entity.User, error)
		Update(context.Context, *entity.User) error
		DeleteById(context.Context, uint) error
		FindAll(context.Context) ([]entity.User, error)
	}

	CapabilityRepo interface {
		Create(context.Context, *entity.Capability) error
		FindById(context.Context, uint) (entity.Capability, error)
		Update(context.Context, *entity.Capability) error
		DeleteById(context.Context, uint) error
		FindAll(context.Context) ([]entity.Capability, error)
	}
)
