package domain

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"query-handler/internal/entity"
	"query-handler/internal/repo/domain/bundle_id"
	"query-handler/internal/repo/domain/capability"
	"query-handler/internal/repo/domain/certificate"
	"query-handler/internal/repo/domain/device"
	"query-handler/internal/repo/domain/integration"
	"query-handler/internal/repo/domain/profile"
	"query-handler/internal/repo/domain/tenant"
	"query-handler/internal/repo/domain/user"
	"query-handler/internal/usecase"
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

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func New(conf PostgresConfig) *PostgresDomainRepo {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.Host,
		conf.User,
		conf.Password,
		conf.Name,
		conf.Port,
	)
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
