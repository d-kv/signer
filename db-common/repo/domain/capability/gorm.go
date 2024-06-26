package capability

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

func (repo *GormRepo) Create(ctx context.Context, capability *entity.Capability) error {
	err := repo.DB.WithContext(ctx).Create(capability).Error
	return err
}

func (repo *GormRepo) FindByBundleIdId(ctx context.Context, ID string) ([]entity.Capability, error) {
	var capabilities []entity.Capability
	err := repo.DB.WithContext(ctx).Where("bundle_id_id = ?", ID).Find(&capabilities).Error
	return capabilities, err
}

func (repo *GormRepo) Update(ctx context.Context, capability *entity.Capability) error {
	err := repo.DB.WithContext(ctx).Save(capability).Error
	return err
}

func (repo *GormRepo) DeleteById(ctx context.Context, ID string) error {
	err := repo.DB.WithContext(ctx).Where("id = ?", ID).Delete(&entity.Capability{}).Error
	return err
}

func (repo *GormRepo) FindAll(ctx context.Context) ([]entity.Capability, error) {
	var capabilities []entity.Capability
	err := repo.DB.WithContext(ctx).Find(&capabilities).Error
	return capabilities, err
}
