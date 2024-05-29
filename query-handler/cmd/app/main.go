package main

import (
	"d-kv/signer/db-common/config"
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
	"os"
)

func main() {
	pgRepo := config.PostgresConfig{
		Host:     os.Getenv("DOMAIN_POSTGRES_HOST"),
		User:     os.Getenv("DOMAIN_POSTGRES_USER"),
		Password: os.Getenv("DOMAIN_POSTGRES_PASSWORD"),
		Name:     os.Getenv("DOMAIN_POSTGRES_DB"),
		Port:     os.Getenv("DOMAIN_POSTGRES_PORT"),
	}
	db, err := gorm.Open(postgres.Open(pgRepo.ToConnectionString()))

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
	log.Println("Query-handler successfully started!")
	if err != nil {
		log.Fatalf("server running error: %s", err.Error())
	}
}
