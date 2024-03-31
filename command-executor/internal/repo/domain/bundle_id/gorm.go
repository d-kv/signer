package bundle_id

import (
	"command-executor/internal/entity"
	"context"
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

func (repo *GormRepo) FindById(ctx context.Context, ID uint) (entity.BundleId, error) {
	var bundleId = entity.BundleId{}
	err := repo.DB.WithContext(ctx).First(&bundleId, ID).Error
	return bundleId, err
}

func (repo *GormRepo) Update(ctx context.Context, bundleId *entity.BundleId) error {
	err := repo.DB.WithContext(ctx).Save(bundleId).Error
	return err
}

func (repo *GormRepo) DeleteById(ctx context.Context, ID uint) error {
	err := repo.DB.WithContext(ctx).Delete(&entity.BundleId{}, ID).Error
	return err
}
