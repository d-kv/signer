package profile

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

func (repo *GormRepo) Create(ctx context.Context, profile *entity.Profile) error {
	err := repo.DB.WithContext(ctx).Create(profile).Error
	return err
}

func (repo *GormRepo) FindById(ctx context.Context, ID string) (entity.Profile, error) {
	var profile = entity.Profile{}
	err := repo.DB.WithContext(ctx).First(&profile, ID).Error
	return profile, err
}

func (repo *GormRepo) Update(ctx context.Context, profile *entity.Profile) error {
	err := repo.DB.WithContext(ctx).Save(profile).Error
	return err
}

func (repo *GormRepo) DeleteById(ctx context.Context, ID string) error {
	err := repo.DB.WithContext(ctx).Delete(&entity.Profile{}, ID).Error
	return err
}

func (repo *GormRepo) FindAll(ctx context.Context) ([]entity.Profile, error) {
	var profiles []entity.Profile
	err := repo.DB.WithContext(ctx).Find(&profiles).Error
	return profiles, err
}
