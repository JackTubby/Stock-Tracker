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
	var lastOffset int

	for {
		updates, err := telegramClient.GetUpdates(lastOffset)
		if err != nil {
			log.Fatal(err)
		}
		
		for _, update := range updates {
			text := update.Message.Text
			chatID := update.Message.Chat.Id
			lastOffset = update.UpdateId + 1
			
			if text == "/start" {
				telegramClient.SendMessage(chatID, "Welcome!")
			}
			if text == "hey" {
				telegramClient.SendMessage(chatID, "You said 'hey'!")
			}
		}
		
		time.Sleep(2 * time.Second)
	}
}
