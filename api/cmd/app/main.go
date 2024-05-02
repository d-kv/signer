package main

import (
	clientservice "d-kv/signer/api/internal/client-service"
	"d-kv/signer/api/pkg/httpserver"
	"d-kv/signer/api/pkg/httpserver/handlers"
	"d-kv/signer/db-common/config"
	"d-kv/signer/db-common/repo/command"
	"log"
)

// @title AppStoreConnect project API
// @version 1.0
// @description API server for AppStoreConnect project

// @host localhost:8080
// @BasePath /v1/{tenantId}/{integrationId}

func main() {
	pgConfig := config.PostgresConfig{Host: "localhost", User: "postgres",
		Password: "postgres", Name: "command_queue", Port: "5432"}
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
