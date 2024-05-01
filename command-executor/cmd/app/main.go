package main

import (
	"context"
	"d-kv/signer/command-executor/internal/services"
	"d-kv/signer/db-common/config"
	"d-kv/signer/db-common/repo/command"
	"d-kv/signer/db-common/repo/domain"
)

func main() {
	//	ctx, cancel := context.WithCancel(context.Background())
	//	go services.Activate(ctx)
	//	time.Sleep(1 * time.Minute)
	//	cancel()
	//	token generation

	pgQueue := config.PostgresConfig{Host: "localhost", User: "postgres",
		Password: "postgres", Name: "command_queue", Port: "5432"}
	pgRepo := domain.PostgresConfig{Host: "localhost", User: "postgres",
		Password: "postgres", Name: "postgres", Port: "5433"}

	queue := command.New(pgQueue)
	repo := domain.New(pgRepo)

	ctx := context.Background()
	services.StartProcessor(ctx, queue, repo)
}
