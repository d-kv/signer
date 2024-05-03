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
		Token:         "LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JR1RBZ0VBTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEJIa3dkd0lCQVFRZ3lFVFZlRmdSQmdCTlpIZlEKUmZ1NUtGMG50T0twUTJwNCtQdjV5bUUwZnoyZ0NnWUlLb1pJemowREFRZWhSQU5DQUFSWktEa2xER2hGdzk1WgovSUNqS1RJUFJ6NU1ucGd5ajZkS29UTkZXdHF3Y1BhSnAzdU1EZitUbG1ZNFJlZitSNmw2QmlnREFSR2t0bFNRCkRRUVNscWJlCi0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0=",
		KeyId:         "CHNPV39B43",
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
