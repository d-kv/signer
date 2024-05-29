package main

import (
	clientservice "d-kv/signer/api/internal/client-service"
	"d-kv/signer/api/pkg/httpserver"
	"d-kv/signer/api/pkg/httpserver/handlers"
	"d-kv/signer/db-common/config"
	"d-kv/signer/db-common/repo/command"
	"log"
	"os"
)

// @title AppStoreConnect project API
// @version 1.0
// @description API server for AppStoreConnect project

// @host localhost:8080
// @BasePath /v1/{tenantId}/{integrationId}

func main() {
	pgConfig := config.PostgresConfig{
		Host:     os.Getenv("COMMAND_QUEUE_POSTGRES_HOST"),
		User:     os.Getenv("COMMAND_QUEUE_POSTGRES_USER"),
		Password: os.Getenv("COMMAND_QUEUE_POSTGRES_PASSWORD"),
		Name:     os.Getenv("COMMAND_QUEUE_POSTGRES_DB"),
		Port:     os.Getenv("COMMAND_QUEUE_POSTGRES_PORT"),
	}
	queue := command.New(pgConfig)
	queryURL := "http://localhost:8081/v1/query"
	services := clientservice.NewService(queue, queryURL)
	handler := handlers.NewHandler(services)
	server := new(httpserver.Server)
	err := server.Run("8080", handler.InitRoutes())
	if err != nil {
		log.Fatalf("server running error: %s", err.Error())
	}
}
