package main

import (
	"command-executor/internal/services"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	jwtToken, err := services.GetJWT([]string{""}, 20*time.Minute)
	if err != nil {
		log.Fatal("Error generate JWT")
		return
	}
	err = os.WriteFile("token.txt", []byte(jwtToken), 0644)
	if err != nil {
		log.Fatal("Error write token")
		return
	}
	fmt.Println(jwtToken)
	fmt.Println("Token was written!")
}
