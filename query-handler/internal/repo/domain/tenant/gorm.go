package tenant

import (
	"context"
	"gorm.io/gorm"
	"query-handler/internal/entity"
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

func (repo *GormRepo) FindById(ctx context.Context, ID uint) (entity.Tenant, error) {
	var tenant = entity.Tenant{}
	err := repo.DB.WithContext(ctx).First(&tenant, ID).Error
	return tenant, err
}

func (repo *GormRepo) Update(ctx context.Context, tenant *entity.Tenant) error {
	err := repo.DB.WithContext(ctx).Save(tenant).Error
	return err
}

func (repo *GormRepo) DeleteById(ctx context.Context, ID uint) error {
	err := repo.DB.WithContext(ctx).Delete(&entity.Tenant{}, ID).Error
	return err
}

func (repo *GormRepo) FindAll(ctx context.Context) ([]entity.Tenant, error) {
	var tenants []entity.Tenant
	err := repo.DB.WithContext(ctx).Find(&tenants).Error
	return tenants, err
}
