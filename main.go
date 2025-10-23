package main

import (
	"fmt"
	"log"
	"os"
	"stock-watch/clients/finnhub"
	"stock-watch/clients/telegram"

	"github.com/joho/godotenv"
)

type Config struct {
	FINNHUB_API_KEY string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	return &Config{
		FINNHUB_API_KEY: mustGetEnv("FINNHUB_API_KEY"),
	}
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is required", key)
	}
	return value
}

func main() {
	config := NewConfig()
	finnhubClient := finnhub.NewClient(config.FINNHUB_API_KEY)
	
	 quote, err := finnhubClient.GetQuote("AAPL")
	 if err != nil {
		 log.Fatal(err) 
	}
	fmt.Println("ingore: ", quote)
	fmt.Println("               ")

	telegramClient := telegram.NewClient()

	getMe, err := telegramClient.GetMe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", getMe)
}

