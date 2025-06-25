package natsx

import (
	"sync"

	"github.com/mixdjoker/chain-notes/internal/config"
)

type natsConfig struct {
	URL string
	// future: TLS, timing, etc
}

var (
	once   sync.Once
	cached *natsConfig
)

func getConfig() *natsConfig {
	once.Do(func() {
		cached = &natsConfig{
			URL: config.GetEnv("NATS_URL", "nats://localhost:4222"),
		}
	})
	return cached
}
