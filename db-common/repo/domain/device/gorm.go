package device

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

func (repo *GormRepo) Create(ctx context.Context, device *entity.Device) error {
	err := repo.DB.WithContext(ctx).Create(device).Error
	return err
}

func (repo *GormRepo) FindById(ctx context.Context, ID string) (entity.Device, error) {
	var device = entity.Device{}
	err := repo.DB.WithContext(ctx).Where("id = ?", ID).First(&device).Error
	return device, err
}

func (repo *GormRepo) Update(ctx context.Context, device *entity.Device) error {
	err := repo.DB.WithContext(ctx).Save(device).Error
	return err
}

func (repo *GormRepo) DeleteById(ctx context.Context, ID string) error {
	err := repo.DB.WithContext(ctx).Where("id = ?", ID).Delete(&entity.Device{}).Error
	return err
}

func (repo *GormRepo) FindAll(ctx context.Context) ([]entity.Device, error) {
	var devices []entity.Device
	err := repo.DB.WithContext(ctx).Find(&devices).Error
	return devices, err
}
