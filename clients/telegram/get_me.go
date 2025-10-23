package telegram

import (
	"io"
	"net/http"
	"fmt"
	"encoding/json"
)

type MeResponse struct {
	Ok bool `json:"ok"`
	Result User `json:"result"`
}

type User struct {
	Id int `json:"id"`
	IsBot bool `json:"is_bot"`
	FirstName string `json:"first_name"`
}

func (c *Client) GetMe() (MeResponse, error) {
	url := c.url + "getMe"
	resp, err := http.Get(url)
	if err != nil {
		return MeResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return MeResponse{}, fmt.Errorf("API error: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return MeResponse{}, err
	}
	
	var result MeResponse
	json.Unmarshal(body, &result)
	return result, nil
}

