package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIKey string
}

func NewConfig() *Config {
	godotenv.Load()

	return &Config{
		APIKey: mustGetEnv("TELEGRAM_KEY"),
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

	url := fmt.Sprintf("https://api.telegram.org/bot%s/getMe", config.APIKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	fmt.Println("Status:", resp.Status)
	fmt.Println("Body:", string(body))
}
