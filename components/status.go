package components

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"runtime"
	"strings"
	"time"
)

type AllocatedMemory struct {
	Total      float64
	GC         float64
	Sys        float64
	Percentage float64
}

func GetAllocatedMemory() AllocatedMemory {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	totalMemory := float64(mem.HeapAlloc) / (1024 * 1024)
	gcMemory := float64(mem.Alloc) / (1024 * 1024)
	sysMemory := float64(mem.Sys) / (1024 * 1024)
	usagePercentage := (totalMemory / sysMemory) * 100

	return AllocatedMemory{
		Total:      totalMemory,
		GC:         gcMemory,
		Sys:        sysMemory,
		Percentage: usagePercentage,
	}
}

func GetUptime(launchTime int64) string {
	currentTime := time.Now().Unix()
	timeIntervalStamp := currentTime - launchTime

	duration := time.Second * time.Duration(timeIntervalStamp)
	days := duration / (time.Hour * 24)
	hours := (duration % (time.Hour * 24)) / time.Hour
	minutes := (duration % time.Hour) / time.Minute
	seconds := duration % time.Minute

	uptime := fmt.Sprintf("%02d:%02d:%02d:%02d", days, hours, minutes, seconds)
	uptime = strings.TrimRight(uptime, "0")

	return uptime
}

func HandleStatusModule(bot *tgbotapi.BotAPI, update tgbotapi.Update, launchTime int64) {
	chatID := update.Message.Chat.ID

	osInfo := runtime.GOOS

	allocateMemory := GetAllocatedMemory()

	botVersion := "v0.0.4"
	botVersionTag := "Alpha"

	statusMessage := fmt.Sprintf(
		"FrostyAssistant - 运行状态\n"+
			"\n"+
			"- OS: %s\n"+
			"- GC: %.2f MB\n"+
			"- Memory: %.2f MB / %.2f MB (%.2f%%)\n"+
			"- Uptime: %s\n"+
			"- Version: %s-%s\n"+
			"\n"+
			"Powered by go-telegram-bot-api\n"+
			"Maintainer: Snowball_233",
		osInfo,
		allocateMemory.GC,
		allocateMemory.Total, allocateMemory.Sys, allocateMemory.Percentage,
		GetUptime(launchTime),
		botVersion, botVersionTag)

	msg := tgbotapi.NewMessage(chatID, statusMessage)
	msg.ReplyToMessageID = update.Message.MessageID
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error occurred, caused by: %v", err)
	}
}
