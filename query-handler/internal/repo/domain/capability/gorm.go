package capability

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

func (repo *GormRepo) Create(ctx context.Context, capability *entity.Capability) error {
	err := repo.DB.WithContext(ctx).Create(capability).Error
	return err
}

func (repo *GormRepo) FindById(ctx context.Context, ID uint) (entity.Capability, error) {
	var capability = entity.Capability{}
	err := repo.DB.WithContext(ctx).First(&capability, ID).Error
	return capability, err
}

func (repo *GormRepo) Update(ctx context.Context, capability *entity.Capability) error {
	err := repo.DB.WithContext(ctx).Save(capability).Error
	return err
}

func (repo *GormRepo) DeleteById(ctx context.Context, ID uint) error {
	err := repo.DB.WithContext(ctx).Delete(&entity.Capability{}, ID).Error
	return err
}
