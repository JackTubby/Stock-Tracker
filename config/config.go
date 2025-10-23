package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	FINNHUB_API_Key string
}

func load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	return &Config {
		FINNHUB_API_Key: mustGetEnv("FINNHUB_API_Key")
	}
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is required", key)
	}
	return value
}

