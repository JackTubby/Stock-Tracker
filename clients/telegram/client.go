package telegram

import (
	"log"
)

type Client struct {
	url string
	token string
}

func NewClient(token string) *Client {
	if token == "" {
		log.Fatalf("Telegram token does not exist!")
		return nil
	}
	fullUrl := "https://api.telegram.org/bot" + token + "/"
	return &Client {
		url: fullUrl,
		token: token,
	}
}
