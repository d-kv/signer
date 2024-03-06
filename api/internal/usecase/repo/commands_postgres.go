package repo

import (
	"api/internal/entity"
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

func (repo *CommandRepo) CreateBundleIdCommand(ctx context.Context, command *entity.CreateBundleId) error {
	result := repo.db.WithContext(ctx).Create(command)
	return result.Error
}

func (repo *CommandRepo) CreateDeviceCommand(ctx context.Context, command *entity.CreateDevice) error {
	result := repo.db.WithContext(ctx).Create(command)
	return result.Error
}

func (repo *CommandRepo) CreateEnableCapabilityTypeCommand(ctx context.Context, command *entity.EnableCapabilityType) error {
	result := repo.db.WithContext(ctx).Create(command)
	return result.Error
}
