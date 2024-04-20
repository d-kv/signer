package main

import (
	"db/repo/domain/bundle_id"
	"db/repo/domain/capability"
	"db/repo/domain/certificate"
	"db/repo/domain/device"
	"db/repo/domain/integration"
	"db/repo/domain/profile"
	"db/repo/domain/tenant"
	"db/repo/domain/user"
	"db/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/url"
	"query-handler/pkg/httpserver"
	"query-handler/pkg/httpserver/handlers"
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
		QueryProcessor: &services.QueryProcessor{
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
