package main

import (
	"log"
	"fmt"
	"time"
	"stock-watch/clients/telegram"
	"stock-watch/clients/finnhub"
	"stock-watch/config"
)

func main() {
	cfg := config.Load()
	
	finnhubClient := finnhub.NewClient(cfg.FinnhubApiKey)
	telegramClient := telegram.NewClient(cfg.TelegramToken)
	var lastOffset int

	for {
		updates, err := telegramClient.GetUpdates(lastOffset)
		if err != nil {
			log.Fatal(err)
		}
		
		for _, update := range updates {
			text := update.Message.Text
			fmt.Println("text: ", text)
			chatID := update.Message.Chat.Id
			lastOffset = update.UpdateId + 1
			
			if text == "/start" {
				telegramClient.SendMessage(chatID, "Welcome!")
			}
			if text == "quote" {
				currentStatus, err := finnhubClient.GetQuote("AAPL")
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(currentStatus)
				// telegramClient.SendMessage(chatID, currentStatus)
			}
		}
		
		time.Sleep(2 * time.Second)
	}
}
