package main

import (
	clientservice "d-kv/signer/api/internal/client-service"
	"d-kv/signer/api/pkg/httpserver"
	"d-kv/signer/api/pkg/httpserver/handlers"
	"log"
)

// @title AppStoreConnect project API
// @version 1.0
// @description API server for AppStoreConnect project

// @host localhost:8080
// @BasePath /v1/{tenantId}/{integrationId}

func main() {
	services := clientservice.NewService()
	handler := handlers.NewHandler(services)
	server := new(httpserver.Server)
	err := server.Run("8080", handler.InitRoutes())
	if err != nil {
		log.Fatalf("server running error: %s", err.Error())
	}
}
