package commitservice

import (
	"context"
)

// Service is a placeholder struct for commitservice package.
type Service struct {
	msg   Messaging
	store Store
	// future: db, logger, metrics, validators
}

type Config struct {
	Messaging Messaging
	Store     Store
}

func New(cfg Config) *Service {
	return &Service{
		msg:   cfg.Messaging,
		store: cfg.Store,
	}
}

func (s *Service) Run(ctx context.Context) error {
	return s.msg.Subscribe(ctx, "chain.commit.submit", "commit-workers", s.handleSubmit)
}
