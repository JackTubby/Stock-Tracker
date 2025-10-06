package finnhub

import (
	"context"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

type Client struct {
	api    *finnhub.DefaultApiService
	ctx    context.Context
	APIKey string
}

func NewClient(apiKey string) *Client {
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", apiKey)

	return &Client{
		api:    finnhub.NewAPIClient(cfg).DefaultApi,
		ctx:    context.Background(),
		APIKey: apiKey,
	}
}
