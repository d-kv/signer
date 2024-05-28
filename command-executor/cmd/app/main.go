package main

import (
	"context"
	"d-kv/signer/command-executor/internal/services"
	ceUsecase "d-kv/signer/command-executor/pkg/usecase"
	"d-kv/signer/db-common/config"
	"d-kv/signer/db-common/entity"
	"d-kv/signer/db-common/repo/command"
	"d-kv/signer/db-common/repo/domain"
	"d-kv/signer/db-common/repo/vault"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()

	vaultConfig := config.VaultConfig{
		Token:   os.Getenv("VAULT_TOKEN"),
		Address: os.Getenv("VAULT_ADDRESS"),
		Timout:  time.Second * 5,
	}
	err, vaultRepo := vault.New(vaultConfig)
	if err != nil {
		log.Fatal(err)
	}
	err = vaultRepo.SaveIntegrationToken(ctx, &entity.IntegrationToken{
		IntegrationId: "iosteam",
		Token:         "NJI32NJ434U3I4U94JN",
		KeyId:         "MF43IMFF3",
	})
	if err != nil {
		log.Fatal(err)
	}

	err, integrationToken := vaultRepo.FindTokenByIntegrationId(ctx, "iosteam")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("integrationToken: %v", integrationToken)

	pgQueue := config.PostgresConfig{
		Host:     os.Getenv("COMMAND_QUEUE_POSTGRES_HOST"),
		User:     os.Getenv("COMMAND_QUEUE_POSTGRES_USER"),
		Password: os.Getenv("COMMAND_QUEUE_POSTGRES_PASSWORD"),
		Name:     os.Getenv("COMMAND_QUEUE_POSTGRES_DB"),
		Port:     os.Getenv("COMMAND_QUEUE_POSTGRES_PORT"),
	}
	pgRepo := config.PostgresConfig{
		Host:     os.Getenv("DOMAIN_POSTGRES_HOST"),
		User:     os.Getenv("DOMAIN_POSTGRES_USER"),
		Password: os.Getenv("DOMAIN_POSTGRES_PASSWORD"),
		Name:     os.Getenv("DOMAIN_POSTGRES_DB"),
		Port:     os.Getenv("DOMAIN_POSTGRES_PORT"),
	}

	queue := command.New(pgQueue)
	repo := domain.New(pgRepo)
	api := services.NewApiService(&http.Client{})
	token := services.NewTokenService(20 * time.Minute)
	db := services.NewDataBaseService(repo)
	service := ceUsecase.NewService(queue, vaultRepo, api, token, db)
	processor := services.NewProcessorService(service)
	log.Println("Server successfully started!")

	processor.StartProcessor(ctx)
}
