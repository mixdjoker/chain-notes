package main

import (
	"context"
	"database/sql"
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

	// Initialize CockroachDB connection
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("database unreachable: %v", err)
	}
	log.Println("[commit-service] connected to database")

	// Initialize application layer
	store := commitservice.NewSQLStore(db)
	app := commitservice.New(commitservice.Config{
		NATS:  nc,
		Store: store,
	})

	// Run message processing loop
	if err := app.Run(ctx); err != nil {
		log.Fatalf("application terminated with error: %v", err)
	}

	log.Println("[commit-service] shutdown complete")
}

type config struct {
	NATSUrl     string
	DatabaseURL string
}

func loadConfig() config {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = "nats://localhost:4222"
	}
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://user:password@localhost:26257/chain_notes?sslmode=disable"
	}
	return config{
		NATSUrl:     url,
		DatabaseURL: databaseURL,
	}
}
