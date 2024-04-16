package certificate

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

func (repo *GormRepo) Create(ctx context.Context, certificate *entity.Certificate) error {
	err := repo.DB.WithContext(ctx).Create(certificate).Error
	return err
}

func (repo *GormRepo) FindById(ctx context.Context, ID uint) (entity.Certificate, error) {
	var certificate = entity.Certificate{}
	err := repo.DB.WithContext(ctx).First(&certificate, ID).Error
	return certificate, err
}

func (repo *GormRepo) Update(ctx context.Context, certificate *entity.Certificate) error {
	err := repo.DB.WithContext(ctx).Save(certificate).Error
	return err
}

func (repo *GormRepo) DeleteById(ctx context.Context, ID uint) error {
	err := repo.DB.WithContext(ctx).Delete(&entity.Certificate{}, ID).Error
	return err
}