package repo

import (
	"command-executor/internal/entity"
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type CommandRepo struct {
	db *gorm.DB
}

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func New(conf PostgresConfig) *CommandRepo {
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
	return &CommandRepo{db: db}
}

func (repo *CommandRepo) FindByStatusBundleIdCommand(ctx context.Context, status entity.Status) []entity.CreateBundleId {
	var commands []entity.CreateBundleId
	repo.db.WithContext(ctx).Where("status = ?", status).Find(&commands)
	return commands
}

func (repo *CommandRepo) FindByStatusDeviceCommand(ctx context.Context, status entity.Status) []entity.CreateDevice {
	var commands []entity.CreateDevice
	repo.db.WithContext(ctx).Where("status = ?", status).Find(&commands)
	return commands
}

func (repo *CommandRepo) FindByStatusEnableCapabilityTypeCommand(ctx context.Context, status entity.Status) []entity.EnableCapabilityType {
	var commands []entity.EnableCapabilityType
	repo.db.WithContext(ctx).Where("status = ?", status).Find(&commands)
	return commands
}

func (repo *CommandRepo) SetStatusByIdBundleIdCommand(ctx context.Context, ID uint, status entity.Status) error {
	err := repo.db.WithContext(ctx).Model(&entity.CreateBundleId{}).Where("id = ?", ID).Update("status", status).Error
	return err
}

func (repo *CommandRepo) SetStatusByIdDeviceCommand(ctx context.Context, ID uint, status entity.Status) error {
	err := repo.db.WithContext(ctx).Model(&entity.CreateDevice{}).Where("id = ?", ID).Update("status", status).Error
	return err
}

func (repo *CommandRepo) SetStatusByIdEnableCapabilityTypeCommand(ctx context.Context, ID uint, status entity.Status) error {
	err := repo.db.WithContext(ctx).Model(&entity.EnableCapabilityType{}).Where("id = ?", ID).Update("status", status).Error
	return err
}
