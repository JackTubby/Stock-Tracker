package finnhub

import finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"

func (c *Client) GetQuote(symbol string) (finnhub.Quote, error) {
	res, _, err := c.api.Quote(c.ctx).Symbol(symbol).Execute()
	if err != nil {
		return finnhub.Quote{}, err
	}
	
	return res, nil
}

