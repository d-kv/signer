package main

import (
	"d-kv/signer/db-common/repo/domain/bundle_id"
	"d-kv/signer/db-common/repo/domain/capability"
	"d-kv/signer/db-common/repo/domain/certificate"
	"d-kv/signer/db-common/repo/domain/device"
	"d-kv/signer/db-common/repo/domain/integration"
	"d-kv/signer/db-common/repo/domain/profile"
	"d-kv/signer/db-common/repo/domain/tenant"
	"d-kv/signer/db-common/repo/domain/user"
	"d-kv/signer/db-common/usecase"
	"d-kv/signer/query-handler/pkg/httpserver"
	"d-kv/signer/query-handler/pkg/httpserver/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/url"
)

func main() {
	dsn := url.URL{
		User:     url.UserPassword("postgres", "postgres"),
		Scheme:   "postgres",
		Host:     "localhost:5433",
		Path:     "postgres",
		RawQuery: "sslmode=disable",
	}
	db, err := gorm.Open(postgres.Open(dsn.String()))

	h := handlers.Handler{
		DomainRepos: &usecase.DomainRepos{
			TenantRepo:      &tenant.GormRepo{DB: db},
			DeviceRepo:      &device.GormRepo{DB: db},
			IntegrationRepo: &integration.GormRepo{DB: db},
			ProfileRepo:     &profile.GormRepo{DB: db},
			BundleIdRepo:    &bundle_id.GormRepo{DB: db},
			CertificateRepo: &certificate.GormRepo{DB: db},
			UserRepo:        &user.GormRepo{DB: db},
			CapabilityRepo:  &capability.GormRepo{DB: db},
		},
	}

	server := new(httpserver.Server)
	err = server.Run("8081", h.InitRoutes())
	if err != nil {
		log.Fatalf("server running error: %s", err.Error())
	}
}
