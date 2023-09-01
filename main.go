package main

import (
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
	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60

	updates, err := bot.GetUpdatesChan(update)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		messageText := update.Message.Text
		chatID := update.Message.Chat.ID

		switch {
		//case messageText == "/start":
		//	welcomeMessage := "欢迎使用机器人！"
		//	tgbotapi.NewMessage(chatID, welcomeMessage).Send(bot)

		case messageText == "/help":
			reply.HandleHelpModule(bot, update)

		default:
			//defaultResponse := "不明白您说的是什么，请使用 /start 获取帮助。"
			//tgbotapi.NewMessage(chatID, defaultResponse).Send(bot)
		}
	}
}
