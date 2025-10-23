package telegram

import (
	"os"
)

type Client struct {
	url string
	token string
}

func NewClient() *Client {
	token := os.Getenv("TELEGRAM_TOKEN")
	fullUrl := "https://api.telegram.org/bot" + token + "/"
	return &Client {
		url: fullUrl,
		token: token,
	}
}
