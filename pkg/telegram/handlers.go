package telegram

import (
	"fmt"
	"math/rand"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart = "start"
	lowLvl       = 10
	midLvl       = 100
	highLvl      = 1000
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
	return nil
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Weclome to my friend. There you can calculate simple equations")
	_, err := b.bot.Send(msg)
	msg = tgbotapi.NewMessage(message.Chat.ID, "Choose your level: 1(Low), 2(Mid), 3(High)")
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don't know this command")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message, updates *tgbotapi.UpdatesChannel) {
	switch message.Text {
	case "1":
		b.questions(lowLvl, message)
	case "2":
		b.questions(midLvl, message)
	case "3":
		b.questions(highLvl, message)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Invalid level!!!")
		b.bot.Send(msg)
	}

}

func (b *Bot) questions(lvl int, message *tgbotapi.Message) {
	correct := 0
	for i := 0; i < 10; i++ {
		x := rand.Intn(lvl)
		y := rand.Intn(lvl)
		q := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("%d + %d", x, y))
		b.bot.Send(q)
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60
		updates, _ := b.bot.GetUpdatesChan(u)
	loop:
		for update := range updates { //wait answer
			answerCh := make(chan string)
			go func() {
				var answer string
				answer = update.Message.Text
				answerCh <- answer
			}()
			select {
			case answer := <-answerCh:
				if answer == strconv.Itoa(x+y) {
					correct++
				}
				break loop
			}
		}
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "You score: "+strconv.Itoa(correct)+" out of 10")
	b.bot.Send(msg)
	msg = tgbotapi.NewMessage(message.Chat.ID, "Choose your level: 1(Low), 2(Mid), 3(High)")
	b.bot.Send(msg)
}
