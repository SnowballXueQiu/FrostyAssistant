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

func GetGreeting() string {
	hourNow := time.Now().Hour()
	greet := ""

	switch {
	case hourNow >= 18:
		greet = "晚上好~"
	case hourNow >= 14:
		greet = "下午好~"
	case hourNow >= 11:
		greet = "中午好~"
	case hourNow >= 9:
		greet = "上午好~"
	case hourNow >= 6:
		greet = "早上好~"
	default:
		greet = "早...早上好?"
	}

	return greet
}

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

func GetPositive(luckPoint int) string {
	index := int(math.Floor(float64(luckPoint) / 100 * float64(len(data.Positive))))

	return fmt.Sprintf(
		"%s(%s)",
		data.Positive[index].Title,
		data.Positive[index].Comment,
	)
}

func GetNegative(luckPoint int) string {
	index := int(math.Floor(float64(luckPoint) / 100 * float64(len(data.Negative))))

	return fmt.Sprintf(
		"%s(%s)",
		data.Negative[index].Title,
		data.Negative[index].Comment,
	)
}

func HandleLuckModule(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	userID := update.Message.From.ID

	currenDate := time.Now().Format("20060102")
	seedCode, _ := strconv.ParseInt(strconv.Itoa(userID)+currenDate, 10, 64)

	luckPoint := GetLuckPoint(seedCode)
	mascotToday := data.Mascot[int(math.Floor(float64(luckPoint)/100*float64(len(data.Mascot))))]

	luckMessage := fmt.Sprintf(
		"%s, 您今天的运势是: %s\n"+
			"\n"+
			"- 点数为: %d\n"+
			"- 宜: %s\n"+
			"- 忌: %s\n"+
			"- 今日吉祥物: %s\n"+
			"\n"+
			"*部分内容来源于 洛谷 , 欢迎您的补充!",
		GetGreeting(),
		GetFortune(luckPoint), luckPoint,
		GetPositive(luckPoint), GetNegative(luckPoint),
		mascotToday)

	msg := tgbotapi.NewMessage(chatID, luckMessage)
	msg.ReplyToMessageID = update.Message.MessageID
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error occurred, caused by: %v", err)
	}
}
