package bundle_id

import (
	"context"
	"d-kv/signer/db-common/entity"
	"gorm.io/gorm"
)

type GormRepo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *GormRepo {
	return &GormRepo{DB: db}
}

func (repo *GormRepo) Create(ctx context.Context, bundleId *entity.BundleId) error {
	err := repo.DB.WithContext(ctx).Create(bundleId).Error
	return err
}

func (repo *GormRepo) FindById(ctx context.Context, ID string) (entity.BundleId, error) {
	var bundleId = entity.BundleId{}
	err := repo.DB.WithContext(ctx).Where("id = ?", ID).First(&bundleId).Error
	return bundleId, err
}

func (repo *GormRepo) FindByIntegrationId(ctx context.Context, ID string) (entity.BundleId, error) {
	var bundleId = entity.BundleId{}
	err := repo.DB.WithContext(ctx).Where("integration_id = ?", ID).First(&bundleId).Error
	return bundleId, err
}

func (repo *GormRepo) Update(ctx context.Context, bundleId *entity.BundleId) error {
	err := repo.DB.WithContext(ctx).Save(bundleId).Error
	return err
}

func (repo *GormRepo) DeleteById(ctx context.Context, ID string) error {
	err := repo.DB.WithContext(ctx).Where("id = ?", ID).Delete(&entity.BundleId{}).Error
	return err
}

func (repo *GormRepo) FindAll(ctx context.Context) ([]entity.BundleId, error) {
	var bundleIds []entity.BundleId
	err := repo.DB.WithContext(ctx).Find(&bundleIds).Error
	return bundleIds, err
}
