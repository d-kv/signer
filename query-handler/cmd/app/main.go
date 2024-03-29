package main

import (
	"log"
	"query-handler/pkg/httpserver"
	"query-handler/pkg/httpserver/handlers"
)

func main() {
	h := new(handlers.Handler)
	server := new(httpserver.Server)
	err := server.Run("8081", h.InitRoutes())
	if err != nil {
		log.Fatalf("server running error: %s", err.Error())
	}
}
