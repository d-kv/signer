package main

import (
	"command-executor/internal/services"
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go services.Activate(ctx)
	time.Sleep(1 * time.Minute)
	cancel()
}
