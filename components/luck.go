package components

import (
	"FrostyAssistant/components/data"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"math"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

func GetLuckPoint(userName string) int {
	// Get current date
	currenDate := time.Now().Format("2006-01-02")

	// Get to ascii code based on userName
	trimmedInput := strings.ReplaceAll(userName, " ", "")
	asciiString := ""
	for _, char := range trimmedInput {
		asciiString += fmt.Sprintf("%d", int(char))
	}

	// Calculate the MD5 hash value of the seedCode and truncate the first 8 characters
	seedCode := strings.TrimSpace(asciiString) + currenDate
	hasher := md5.New()
	hasher.Write([]byte(seedCode))
	seedMd5 := hex.EncodeToString(hasher.Sum(nil))[:8]

	// Convert seedMd5 to BigInt type
	seedBigInt, _ := new(big.Int).SetString(seedMd5, 16)
	seed := seedBigInt.Int64()

	// Define Algorithm Constants
	m := int64(4294967296)
	a := int64(1103515245)
	c := int64(12345)

	// Calculate luckPoint
	luckPoint := (float64((a*seed+c)%m) / float64(m-1)) * 100

	return int(luckPoint)
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

func GetPositive() string {
	rand.Seed(time.Now().UnixNano())
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

func GetNegative() string {
	rand.Seed(time.Now().UnixNano())
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
	userName := update.Message.From.UserName

	luckPoint := GetLuckPoint(userName)

	luckMessage := fmt.Sprintf("您今天的运势是: %s\n"+
		"- 点数为: %d\n"+
		"- 宜: %s\n"+
		"- 忌: %s\n"+
		"*部分内容来源于 **洛谷**",
		GetFortune(luckPoint), luckPoint,
		GetPositive(), GetNegative())

	msg := tgbotapi.NewMessage(chatID, luckMessage)
	msg.ReplyToMessageID = update.Message.MessageID
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error occurred, caused by: %v", err)
	}
}
