package certificate

import (
	"context"
	"d-kv/signer/query-handler/internal/entity"
	"gorm.io/gorm"
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

func (repo *GormRepo) FindById(ctx context.Context, ID string) (entity.Certificate, error) {
	var certificate = entity.Certificate{}
	err := repo.DB.WithContext(ctx).First(&certificate, ID).Error
	return certificate, err
}

func (repo *GormRepo) Update(ctx context.Context, certificate *entity.Certificate) error {
	err := repo.DB.WithContext(ctx).Save(certificate).Error
	return err
}

func (repo *GormRepo) DeleteById(ctx context.Context, ID string) error {
	err := repo.DB.WithContext(ctx).Delete(&entity.Certificate{}, ID).Error
	return err
}

func (repo *GormRepo) FindAll(ctx context.Context) ([]entity.Certificate, error) {
	var certificates []entity.Certificate
	err := repo.DB.WithContext(ctx).Find(&certificates).Error
	return certificates, err
}
