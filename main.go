package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5660364356:AAFQE51uUghxdGcJ7fL0FCV7cy5Su6pZjDI")

	if err != nil {
		log.Panic("Cannot initialize telegram bot")
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			msg_content := ""

			if update.Message.Text == "/start" {
				msg_content = fmt.Sprintf("Welcome to whatsapp helper you could call any number without saving it by writing the phone number ex: 012xxxxxxxx")
			} else {
				msg_content = fmt.Sprintf("https://wa.me/+2%s", update.Message.Text)
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msg_content)

			bot.Send(msg)
		}
	}
}
