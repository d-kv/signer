package domain

import (
	"d-kv/signer/db-common/config"
	"d-kv/signer/db-common/entity"
	"d-kv/signer/db-common/repo/domain/bundle_id"
	"d-kv/signer/db-common/repo/domain/capability"
	"d-kv/signer/db-common/repo/domain/certificate"
	"d-kv/signer/db-common/repo/domain/device"
	"d-kv/signer/db-common/repo/domain/integration"
	"d-kv/signer/db-common/repo/domain/profile"
	"d-kv/signer/db-common/repo/domain/tenant"
	"d-kv/signer/db-common/repo/domain/user"
	"d-kv/signer/db-common/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type PostgresDomainRepo struct {
	TenantRepo      usecase.TenantRepo
	DeviceRepo      usecase.DeviceRepo
	IntegrationRepo usecase.IntegrationRepo
	ProfileRepo     usecase.ProfileRepo
	BundleIdRepo    usecase.BundleIdRepo
	CertificateRepo usecase.CertificateRepo
	UserRepo        usecase.UserRepo
	CapabilityRepo  usecase.CapabilityRepo
}

func New(conf config.PostgresConfig) *PostgresDomainRepo {
	dsn := conf.ToConnectionString()
	return NewFromDsn(dsn)
}

func NewFromDsn(dsn string) *PostgresDomainRepo {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error open gorm connection wti db", err)
	}
	err = db.AutoMigrate(
		&entity.Tenant{},
		&entity.Device{},
		&entity.Integration{},
		&entity.Profile{},
		&entity.BundleId{},
		&entity.Certificate{},
		&entity.User{},
		&entity.Capability{},
	)
	if err != nil {
		log.Fatal(err)
	}
	return &PostgresDomainRepo{
		TenantRepo:      &tenant.GormRepo{DB: db},
		DeviceRepo:      &device.GormRepo{DB: db},
		IntegrationRepo: &integration.GormRepo{DB: db},
		ProfileRepo:     &profile.GormRepo{DB: db},
		BundleIdRepo:    &bundle_id.GormRepo{DB: db},
		CertificateRepo: &certificate.GormRepo{DB: db},
		UserRepo:        &user.GormRepo{DB: db},
		CapabilityRepo:  &capability.GormRepo{DB: db},
	}
}
