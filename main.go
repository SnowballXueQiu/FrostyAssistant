package main

import (
	"FrostyAssistant/components"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	// Set up Telegram API
	bot, err := tgbotapi.NewBotAPI("YOUR_BOT_TOKEN")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	// Set up update config, and launch Bot
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

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
