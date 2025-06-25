package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/mixdjoker/chain-notes/internal/app/commitservice"
	"github.com/mixdjoker/chain-notes/internal/config"
	"github.com/mixdjoker/chain-notes/internal/dbx"
	"github.com/mixdjoker/chain-notes/internal/infra/messaging"
	"github.com/mixdjoker/chain-notes/internal/infra/natsx"
)

func main() {
	log.Println("[commit-service] starting...")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := config.InitEnv(); err != nil {
		log.Fatalf("[commit-service] failed to load .env: %v", err)
	}

	// Initialize NATS connection
	nc, err := natsx.Connect()
	if err != nil {
		log.Println("[commit-service] started without NATS connection")
	}
	defer nc.Drain()

	// Initialize CockroachDB connection
	db := dbx.Get()
	defer db.Close()

	// Create messaging adapter
	natsClient := messaging.New(nc)

	// Create SQL store for commit service
	store := commitservice.NewSQLStore(db)

	// Initialize application layer
	app := commitservice.New(commitservice.Config{
		Messaging: natsClient,
		Store:     store,
	})

	// Run message processing loop
	if err := app.Run(ctx); err != nil {
		log.Fatalf("application terminated with error: %v", err)
	}

	log.Println("[commit-service] shutdown complete")
}
