package finnhub

import finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"

func (c *Client) GetMarketStatus(exchange string) (finnhub.MarketStatus, error) {
	res, _, err := c.api.MarketStatus(c.ctx).Exchange(exchange).Execute()
	if err != nil {
		return finnhub.MarketStatus{}, err
	}
	return res, nil
}
