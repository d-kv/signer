package user

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

func (repo *GormRepo) Create(ctx context.Context, user *entity.User) error {
	err := repo.DB.WithContext(ctx).Create(user).Error
	return err
}

func (repo *GormRepo) FindById(ctx context.Context, ID uint) (entity.User, error) {
	var user = entity.User{}
	err := repo.DB.WithContext(ctx).First(&user, ID).Error
	return user, err
}

func (repo *GormRepo) Update(ctx context.Context, user *entity.User) error {
	err := repo.DB.WithContext(ctx).Save(user).Error
	return err
}

func (repo *GormRepo) DeleteById(ctx context.Context, ID uint) error {
	err := repo.DB.WithContext(ctx).Delete(&entity.User{}, ID).Error
	return err
}
