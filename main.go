package main

import (
	"FrostyAssistant/components"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Set up Telegram API
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_HTTP_API"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	// Set up update config, and launch Bot
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60


	for update := range updates {
		if update.Message == nil {
			continue
		}

		messageText := update.Message.Text
		chatID := update.Message.Chat.ID

		switch {
		case messageText == "/help":
			components.HandleHelpModule(bot, update)

		case messageText == "/luck":
			components.HandleLuckModule(bot, update)

		default:
			defaultResponse := "本喵真是看不懂一点, 请使用 /help 获取帮助!"
			tgbotapi.NewMessage(chatID, defaultResponse).Send()
		}
	}
}
