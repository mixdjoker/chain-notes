package dbx

import (
	"sync"

	"github.com/mixdjoker/chain-notes/internal/config"
)

type dbConfig struct {
	URL string
}

var (
	once   sync.Once
	cashed *dbConfig
)

func getConfig() *dbConfig {
	once.Do(func() {
		cashed = &dbConfig{
			URL: config.GetEnv("DATABASE_URL", ""),
		}
	})
	return cashed
}
