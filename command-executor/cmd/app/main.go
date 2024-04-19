package main

import (
	"command-executor/internal/config"
	"command-executor/internal/repo/command"
	"command-executor/internal/repo/domain"
	"command-executor/internal/services"
	"context"
)

func main() {
	//	ctx, cancel := context.WithCancel(context.Background())
	//	go services.Activate(ctx)
	//	time.Sleep(1 * time.Minute)
	//	cancel()
	//	token generation
	pgConfig := config.PostgresConfig{Host: "localhost", User: "postgres",
		Password: "postgres", Name: "command_queue", Port: "54321"}
	queue := command.New(pgConfig)
	pgConfig = config.PostgresConfig{Host: "localhost", User: "postgres",
		Password: "postgres", Name: "postgres", Port: "5433"}
	repo := domain.New(pgConfig)
	ctx := context.Background()
	services.StartProcessor(ctx, queue, repo)
}
