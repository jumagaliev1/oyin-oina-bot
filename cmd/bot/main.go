package main

import (
	"github.com/jumagaliev1/telegrambot/internal/config"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jumagaliev1/telegrambot/pkg/telegram"
)

func main() {
	cfg := config.GetConfig()
	bot, err := tgbotapi.NewBotAPI(cfg.Bot.Token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
