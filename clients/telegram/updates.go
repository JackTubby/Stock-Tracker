package telegram

import (
	"fmt"
	"io"
	"encoding/json"
	"net/http"
)

type UpdatesResponse struct {
	Ok     bool   `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageId int    `json:"message_id"`
	Chat      Chat   `json:"chat"`
	Text      string `json:"text"`
}

type Chat struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
}

func (c *Client) GetUpdates() ([]Update, error) {
	url := c.url + "getUpdates"
	fmt.Println("URL: ", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error 'Updates': %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	var result UpdatesResponse
	json.Unmarshal(body, &result)
	return result.Result, nil
}

