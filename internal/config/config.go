package config

import (
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() error {
	//TODO: add custom .env file path
	return godotenv.Overload()
}

func GetEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}

	return val
}
