package dbx

import (
	"database/sql"
	"log"
	"sync"
)

var (
	dbOnce sync.Once
	db     *sql.DB
	dbErr  error
)

// Get возвращает лениво проинициализированное соединение с базой
func Get() *sql.DB {
	dbOnce.Do(func() {
		cfg := getConfig()
		db, dbErr = sql.Open("postgres", cfg.URL)
		if dbErr != nil {
			log.Fatalf("[dbx] failed to open database: %v", dbErr)
		}

		if err := db.Ping(); err != nil {
			log.Fatalf("[dbx] database unreachable: %v", err)
		}

		log.Println("[dbx] connected to database")
	})
	return db
}
