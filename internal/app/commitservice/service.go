package commitservice

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
)

// Service is a placeholder struct for commitservice package.
type Service struct {
	nc    *nats.Conn
	store Store
	// future: db, logger, metrics, validators
}

type Config struct {
	NATS  *nats.Conn
	Store Store
}

func New(cfg Config) *Service {
	return &Service{
		nc:    cfg.NATS,
		store: cfg.Store,
	}
}

func (s *Service) Run(ctx context.Context) error {
	sub, err := s.nc.QueueSubscribe("chain.commit.submit", "commit-workers", s.handleSubmit)
	if err != nil {
		return err
	}
	log.Println("[commit-service] subscribed to chain.commit.submit")

	<-ctx.Done()
	_ = sub.Unsubscribe()
	return nil
}
