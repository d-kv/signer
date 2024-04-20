package main

import (
	"context"
	"d-kv/signer/command-executor/internal/config"
	"d-kv/signer/command-executor/internal/repo/command"
	"d-kv/signer/command-executor/internal/services"
)

func main() {
	//	ctx, cancel := context.WithCancel(context.Background())
	//	go services.Activate(ctx)
	//	time.Sleep(1 * time.Minute)
	//	cancel()
	//	token generation
	pgConfig := config.PostgresConfig{Host: "localhost", User: "postgres",
		Password: "postgres", Name: "command_queue", Port: "54321"}
	repo := command.New(pgConfig)
	ctx := context.Background()
	services.StartProcessor(ctx, repo)
}
