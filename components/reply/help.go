package components

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func HandleHelpModule(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	messageText := update.Message.Text
	chatID := update.Message.Chat.ID

	helpMessage := `
FrostyAssistant-Bot 使用帮助
/help - 显示帮助菜单
...
`
	msg := tgbotapi.NewMessage(chatID, helpMessage)
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error occurred, caused by: %v", err)
	}
}
