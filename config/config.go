package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	FinnhubApiKey string
	TelegramToken string
}

func Load() *Config {
	godotenv.Load()
	return &Config{
		FinnhubApiKey:  os.Getenv("FINNHUB_API_KEY"),
		TelegramToken:  os.Getenv("TELEGRAM_TOKEN"),
	}
}

