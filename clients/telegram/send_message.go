package telegram
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SendMessageResponse struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

func (c *Client) SendMessage(chatID int, text string) (SendMessageResponse, error) {
	url := c.url + "sendMessage"
	
	payload := map[string]interface{}{
		"chat_id": chatID,
		"text":    text,
	}
	
	body, _ := json.Marshal(payload)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return SendMessageResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return SendMessageResponse{}, fmt.Errorf("API error 'SendMessage': %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	
	respBody, _ := io.ReadAll(resp.Body)
	var result SendMessageResponse
	json.Unmarshal(respBody, &result)
	return result, nil
}

