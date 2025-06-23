package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

func main() {
	log.Println("[commit-service] starting...")

	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = nats.DefaultURL
	}

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("failed to connect to NATS: %v", err)
	}
	defer nc.Drain()

	log.Println("[commit-service] connected to NATS")

	sub, err := nc.QueueSubscribe("chain.commit.submit", "commit-workers", handleCommitSubmit)
	if err != nil {
		log.Fatalf("failed to subscribe: %v", err)
	}
	defer sub.Unsubscribe()

	// Wait for termination
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	log.Println("[commit-service] shutting down")
}

func handleCommitSubmit(msg *nats.Msg) {
	log.Printf("[commit-service] received message: %s", string(msg.Data))

	// TODO: parse JSON, verify signature, validate parent, write to DB

	// simulate successful commit
	response := []byte(`{"status":"ok"}`)
	_ = msg.Respond(response)
}
