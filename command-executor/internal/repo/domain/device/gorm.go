package device

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

func (repo *GormRepo) Create(ctx context.Context, device *entity.Device) error {
	err := repo.DB.WithContext(ctx).Create(device).Error
	return err
}

func (repo *GormRepo) FindById(ctx context.Context, ID uint) (entity.Device, error) {
	var device = entity.Device{}
	err := repo.DB.WithContext(ctx).First(&device, ID).Error
	return device, err
}

func (repo *GormRepo) Update(ctx context.Context, device *entity.Device) error {
	err := repo.DB.WithContext(ctx).Save(device).Error
	return err
}

func (repo *GormRepo) DeleteById(ctx context.Context, ID uint) error {
	err := repo.DB.WithContext(ctx).Delete(&entity.Device{}, ID).Error
	return err
}
