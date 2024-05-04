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
		require.Equal(t, tenant, foundTenant)

		err = repo.TenantRepo.DeleteById(ctx, "tenant_id")
		require.NoError(t, err)
	})

	t.Run("TestDeviceRepo", func(t *testing.T) {
		// Здесь вы можете добавить тесты для DeviceRepo
	})

	t.Run("TestIntegrationRepo", func(t *testing.T) {
		// Здесь вы можете добавить тесты для IntegrationRepo
	})

	t.Run("TestProfileRepo", func(t *testing.T) {
		// Здесь вы можете добавить тесты для ProfileRepo
	})

	t.Run("TestBundleIdRepo", func(t *testing.T) {
		// Здесь вы можете добавить тесты для BundleIdRepo
	})

	t.Run("TestCertificateRepo", func(t *testing.T) {
		// Здесь вы можете добавить тесты для CertificateRepo
	})

	t.Run("TestUserRepo", func(t *testing.T) {
		// Здесь вы можете добавить тесты для UserRepo
	})

	t.Run("TestCapabilityRepo", func(t *testing.T) {
		// Здесь вы можете добавить тесты для CapabilityRepo
	})
}
