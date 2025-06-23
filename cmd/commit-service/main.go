package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mixdjoker/chain-notes/internal/app/commitservice"
	"github.com/mixdjoker/chain-notes/internal/infra/natsx"
)

func main() {
	log.Println("[commit-service] starting...")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := loadConfig()

	// Initialize NATS connection
	nc, err := natsx.Connect(cfg.NATSUrl)
	if err != nil {
		log.Fatalf("failed to connect to NATS: %v", err)
	}
	defer nc.Drain()

	log.Println("[commit-service] connected to NATS")

	// Initialize application layer
	app := commitservice.New(commitservice.Config{
		NATS: nc,
		// future: DB, Logger, Metrics...
	})

	// Run message processing loop
	if err := app.Run(ctx); err != nil {
		log.Fatalf("application terminated with error: %v", err)
	}

	log.Println("[commit-service] shutdown complete")
}

type config struct {
	NATSUrl string
}

func loadConfig() config {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = "nats://localhost:4222"
	}
	return config{NATSUrl: url}
}
