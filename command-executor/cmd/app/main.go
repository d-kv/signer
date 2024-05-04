package main

import (
	"context"
	"d-kv/signer/command-executor/internal/services"
	"d-kv/signer/db-common/config"
	"d-kv/signer/db-common/entity"
	"d-kv/signer/db-common/repo/command"
	"d-kv/signer/db-common/repo/domain"
	"d-kv/signer/db-common/repo/vault"
	"log"
	"time"
)

func main() {
	//	ctx, cancel := context.WithCancel(context.Background())
	//	go services.Activate(ctx)
	//	time.Sleep(1 * time.Minute)
	//	cancel()
	//	token generation
	ctx := context.Background()

	vaultConfig := config.VaultConfig{
		Token:   "my-token",
		Address: "http://localhost:8200",
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

	pgQueue := config.PostgresConfig{Host: "localhost", User: "postgres",
		Password: "postgres", Name: "command_queue", Port: "5432"}
	pgRepo := domain.PostgresConfig{Host: "localhost", User: "postgres",
		Password: "postgres", Name: "postgres", Port: "5433"}

	queue := command.New(pgQueue)
	repo := domain.New(pgRepo)
	service := services.NewProcessorService(queue, repo, vaultRepo)

	log.Println("Server successfully started!")

	service.StartProcessor(ctx)
}
