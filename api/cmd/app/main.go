package main

import (
	"api/pkg/httpserver"
	"api/pkg/httpserver/handlers"
	"log"
)

func main() {
	handlers := new(handlers.Handler)
	server := new(httpserver.Server)
	err := server.Run("8000", handlers.InitRoutes())
	if err != nil {
		log.Fatalf("server running error: %s", err.Error())
	}
}
