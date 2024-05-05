package domain

import (
	"context"
	"d-kv/signer/db-common/config"
	"d-kv/signer/db-common/entity"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostgresDomainRepo(t *testing.T) {
	ctx := context.Background()

	conf := config.PostgresConfig{
		Host:     "localhost",
		User:     "postgres",
		Password: "postgres",
		Name:     "postgres",
		Port:     "5433",
	}

	repo := New(conf)

	t.Run("TestTenantRepo", func(t *testing.T) {
		tenant := entity.Tenant{
			ID:   "tenant_id",
			Name: "some_tenant_name",
		}

		err := repo.TenantRepo.Create(ctx, &tenant)
		require.NoError(t, err)

		allTenants, err := repo.TenantRepo.FindAll(ctx)
		require.NoError(t, err)
		require.Len(t, allTenants, 1)
		require.ElementsMatch(t, []entity.Tenant{tenant}, allTenants)

		foundTenant, err := repo.TenantRepo.FindById(ctx, "tenant_id")
		require.NoError(t, err)
		require.NotNil(t, foundTenant)
		require.Equal(t, tenant, foundTenant)
	})

	t.Run("TestUserRepo", func(t *testing.T) {
		user := entity.User{
			ID:   "userId",
			Name: "Maks",
		}
		err := repo.UserRepo.Create(ctx, &user)
		require.NoError(t, err)
	})

	t.Run("TestIntegrationRepo", func(t *testing.T) {
		integration := entity.Integration{
			ID:       "integrationId",
			IssuerId: "issuerId",
			TeamId:   "teamId",
			TenantId: "tenant_id",
		}
		err := repo.IntegrationRepo.Create(ctx, &integration)
		require.NoError(t, err)
	})

	t.Run("TestDeviceRepo", func(t *testing.T) {
		integration := entity.Integration{
			ID:       "integrationId",
			IssuerId: "issuerId",
			TeamId:   "teamId",
			TenantId: "tenant_id",
		}
		device := entity.Device{
			UDID:         "UDID",
			Name:         "deviceName",
			Platform:     string(entity.Ios),
			Integrations: []entity.Integration{integration},
			UserID:       "userId",
		}
		err := repo.DeviceRepo.Create(ctx, &device)
		require.NoError(t, err)

		foundDevices, err := repo.DeviceRepo.FindByIntegrationId(ctx, "integrationId")
		require.NoError(t, err)
		require.Len(t, foundDevices, 1)
		require.Equal(t, device.UDID, foundDevices[0].UDID)
	})

	t.Run("TestProfileRepo", func(t *testing.T) {
		// Здесь вы можете добавить тесты для ProfileRepo
	})

	t.Run("TestBundleIdRepo", func(t *testing.T) {
		bundleId := entity.BundleId{
			ID:            "bundleID",
			Identifier:    "bundleIdentifier",
			Name:          "bundleName",
			IntegrationId: "integrationId",
		}
		err := repo.BundleIdRepo.Create(ctx, &bundleId)
		require.NoError(t, err)

		foundBundleId, err := repo.BundleIdRepo.FindById(ctx, bundleId.ID)
		require.NoError(t, err)
		require.Equal(t, bundleId, foundBundleId)

		foundBundleId, err = repo.BundleIdRepo.FindByIntegrationId(ctx, "integrationId")
		require.NoError(t, err)
		require.Equal(t, bundleId, foundBundleId)

		bundleId = entity.BundleId{
			ID:            "bundleID.2",
			Identifier:    "bundleIdentifier",
			Name:          "bundleName",
			IntegrationId: "invalidIntegrationId",
		}
		err = repo.BundleIdRepo.Create(ctx, &bundleId)
		require.Error(t, err)
	})

	t.Run("TestCertificateRepo", func(t *testing.T) {
		// Здесь вы можете добавить тесты для CertificateRepo
	})

	t.Run("TestCapabilityRepo", func(t *testing.T) {
		capability := entity.Capability{
			Type:       string(entity.InAppPurchase),
			BundleIdId: "bundleID",
		}
		err := repo.CapabilityRepo.Create(ctx, &capability)
		require.NoError(t, err)

		foundCapabilities, err := repo.CapabilityRepo.FindByBundleIdId(ctx, "bundleID")
		require.NoError(t, err)
		require.Equal(t, string(entity.InAppPurchase), foundCapabilities[0].Type)
	})
}
