package main

import (
	"log"
	"time"
	"stock-watch/clients/telegram"
	"stock-watch/config"
)

func main() {
	cfg := config.Load()
	
	telegramClient := telegram.NewClient(cfg.TelegramToken)
	
	for {
		updates, err := telegramClient.GetUpdates()
		if err != nil {
			log.Fatal(err)
		}
		
		for _, update := range updates {
			text := update.Message.Text
			chatID := update.Message.Chat.Id
			
			if text == "/start" {
				telegramClient.SendMessage(chatID, "Welcome!")
			}
		}
		
		time.Sleep(2 * time.Second)
	}
}
