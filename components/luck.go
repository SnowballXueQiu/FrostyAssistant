package components

import (
	"FrostyAssistant/components/data"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func GetLuckPoint(seedCode int64) int {
	rand.Seed(seedCode)

	return rand.Intn(100)
}

func GetFortune(luckPoint int) string {
	switch {
	case luckPoint >= 80:
		return "大吉"
	case luckPoint >= 60:
		return "吉"
	case luckPoint >= 40:
		return "中平"
	case luckPoint >= 20:
		return "凶"
	default:
		return "大凶"
	}
}

func GetPositive(seedCode int64) string {
	rand.Seed(seedCode)
	randomToday := rand.Float64()

	positiveTitles := make([]string, 0, len(data.Positive))
	positiveComments := make(map[string]string)
	for title, comment := range data.Positive {
		positiveTitles = append(positiveTitles, title)
		positiveComments[title] = comment
	}

	positiveTitle := positiveTitles[int(math.Floor(randomToday*float64(len(positiveTitles))))]
	positiveComment := positiveComments[positiveTitle]

	return fmt.Sprintf("%s(%s)", positiveTitle, positiveComment)
}

func GetNegative(seedCode int64) string {
	rand.Seed(seedCode)
	randomToday := rand.Float64()

	negativeTitles := make([]string, 0, len(data.Negative))
	negativeComments := make(map[string]string)
	for title, comment := range data.Negative {
		negativeTitles = append(negativeTitles, title)
		negativeComments[title] = comment
	}

	negativeTitle := negativeTitles[int(math.Floor(randomToday*float64(len(negativeTitles))))]
	negativeComment := negativeComments[negativeTitle]

	return fmt.Sprintf("%s(%s)", negativeTitle, negativeComment)
}

func HandleLuckModule(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	userID := update.Message.From.ID

	currenDate := time.Now().Format("20060102")
	seedCode, _ := strconv.ParseInt(strconv.Itoa(userID)+currenDate, 10, 64)

	luckPoint := GetLuckPoint(seedCode)

	luckMessage := fmt.Sprintf("您今天的运势是: %s\n"+
		"- 点数为: %d\n"+
		"- 宜: %s\n"+
		"- 忌: %s\n"+
		"*部分内容来源于 **洛谷**",
		GetFortune(luckPoint), luckPoint,
		GetPositive(seedCode), GetNegative(seedCode))

	msg := tgbotapi.NewMessage(chatID, luckMessage)
	msg.ReplyToMessageID = update.Message.MessageID
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error occurred, caused by: %v", err)
	}
}
