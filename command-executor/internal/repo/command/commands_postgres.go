package command

import (
	"command-executor/internal/config"
	"command-executor/internal/entity"
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Repo struct {
	db *gorm.DB
}

func New(conf config.PostgresConfig) *Repo {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.Host,
		conf.User,
		conf.Password,
		conf.Name,
		conf.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error open gorm connection wti db", err)
	}
	err = db.AutoMigrate(
		&entity.EnableCapabilityType{},
		&entity.CreateDevice{},
		&entity.CreateBundleId{},
	)
	if err != nil {
		log.Fatal(err)
	}
	return &Repo{db: db}
}

func (repo *Repo) FindByStatusBundleIdCommand(ctx context.Context, status entity.Status) []entity.CreateBundleId {
	var commands []entity.CreateBundleId
	repo.db.WithContext(ctx).Where("status = ?", status).Find(&commands)
	return commands
}

func (repo *Repo) FindByStatusDeviceCommand(ctx context.Context, status entity.Status) []entity.CreateDevice {
	var commands []entity.CreateDevice
	repo.db.WithContext(ctx).Where("status = ?", status).Find(&commands)
	return commands
}

func (repo *Repo) FindByStatusEnableCapabilityTypeCommand(ctx context.Context, status entity.Status) []entity.EnableCapabilityType {
	var commands []entity.EnableCapabilityType
	repo.db.WithContext(ctx).Where("status = ?", status).Find(&commands)
	return commands
}

func (repo *Repo) SetStatusByIdBundleIdCommand(ctx context.Context, ID uint, status entity.Status) error {
	err := repo.db.WithContext(ctx).Model(&entity.CreateBundleId{}).Where("id = ?", ID).Update("status", status).Error
	return err
}

func (repo *Repo) SetStatusByIdDeviceCommand(ctx context.Context, ID uint, status entity.Status) error {
	err := repo.db.WithContext(ctx).Model(&entity.CreateDevice{}).Where("id = ?", ID).Update("status", status).Error
	return err
}

func (repo *Repo) SetStatusByIdEnableCapabilityTypeCommand(ctx context.Context, ID uint, status entity.Status) error {
	err := repo.db.WithContext(ctx).Model(&entity.EnableCapabilityType{}).Where("id = ?", ID).Update("status", status).Error
	return err
}
