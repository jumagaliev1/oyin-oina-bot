package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jumagaliev1/telegrambot/pkg/telegram"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5589607135:AAG7osV5IEz-VFdBkHE3BtCfmOpquX4yp8I")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
