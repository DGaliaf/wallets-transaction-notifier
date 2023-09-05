package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (b Bot) handleUpdates(u tgbotapi.UpdateConfig) {
	for update := range b.bot.GetUpdatesChan(u) {
		if update.Message != nil {
			if update.Message.IsCommand() {
				b.handleCommands(update)
			} else {
				b.handleMessages(update)
			}
		}
	}
}

func (b Bot) handleMessages(update tgbotapi.Update) {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID

	b.bot.Send(msg)
}

func (b Bot) handleCommands(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		b.processStart(update)
	case "track":
		b.processTrack(update)
	case "profile":
		b.processProfile(update)
		//default:
		//	msg.Text = "I don't know that command"
	}
}
