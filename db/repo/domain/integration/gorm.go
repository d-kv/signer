package integration

import (
	"context"
	"db/entity"
	"gorm.io/gorm"
)

type GormRepo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *GormRepo {
	return &GormRepo{DB: db}
}

func (repo *GormRepo) Create(ctx context.Context, integration *entity.Integration) error {
	err := repo.DB.WithContext(ctx).Create(integration).Error
	return err
}

func (repo *GormRepo) FindById(ctx context.Context, ID string) (entity.Integration, error) {
	var integration = entity.Integration{}
	err := repo.DB.WithContext(ctx).First(&integration, ID).Error
	return integration, err
}

func (repo *GormRepo) Update(ctx context.Context, integration *entity.Integration) error {
	err := repo.DB.WithContext(ctx).Save(integration).Error
	return err
}

func (repo *GormRepo) DeleteById(ctx context.Context, ID string) error {
	err := repo.DB.WithContext(ctx).Delete(&entity.Integration{}, ID).Error
	return err
}

func (repo *GormRepo) FindAll(ctx context.Context) ([]entity.Integration, error) {
	var integrations []entity.Integration
	err := repo.DB.WithContext(ctx).Find(&integrations).Error
	return integrations, err
}
