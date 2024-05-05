package tenant

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

func (repo *GormRepo) Create(ctx context.Context, tenant *entity.Tenant) error {
	err := repo.DB.WithContext(ctx).Create(tenant).Error
	return err
}

func (repo *GormRepo) FindById(ctx context.Context, ID string) (entity.Tenant, error) {
	var tenant = entity.Tenant{}
	err := repo.DB.WithContext(ctx).Where("id = ?", ID).First(&tenant).Error
	return tenant, err
}

func (repo *GormRepo) Update(ctx context.Context, tenant *entity.Tenant) error {
	err := repo.DB.WithContext(ctx).Save(tenant).Error
	return err
}

func (repo *GormRepo) DeleteById(ctx context.Context, ID string) error {
	err := repo.DB.WithContext(ctx).Where("id = ?", ID).Delete(&entity.Tenant{}).Error
	return err
}

func (repo *GormRepo) FindAll(ctx context.Context) ([]entity.Tenant, error) {
	var tenants []entity.Tenant
	err := repo.DB.WithContext(ctx).Find(&tenants).Error
	return tenants, err
}
