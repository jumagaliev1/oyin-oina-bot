package telegram

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {

	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}
		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}
		correct := 0
		for i := 0; i < 10; i++ {
			x := rand.Intn(10)
			y := rand.Intn(10)
			q := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%d + %d", x, y))
			b.bot.Send(q)
			for updateQ := range updates {
				if updateQ.Message.Text == strconv.Itoa(x+y) {
					correct++
					break
				} else {
					msg := tgbotapi.NewMessage(updateQ.Message.Chat.ID, fmt.Sprintf("Wrong answer"))
					b.bot.Send(msg)
					break
				}
			}
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("You deserverd correct %d out 10", correct))
		b.bot.Send(msg)
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint("Choose your level: 1(Low), 2(Mid), 3(High)"))
		b.bot.Send(msg)
	}
}
func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
