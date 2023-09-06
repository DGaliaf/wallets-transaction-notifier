package bot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (b Bot) handleUpdates(ctx context.Context, u tgbotapi.UpdateConfig) {
	for update := range b.bot.GetUpdatesChan(u) {
		if update.Message != nil {
			if update.Message.IsCommand() {
				b.handleCommands(ctx, update)
			} else {
				b.handleMessages(ctx, update)
			}
		}
	}
}

func (b Bot) handleMessages(ctx context.Context, update tgbotapi.Update) {
	if b.token != nil {
		if *b.token == "/track" {
			b.processAddWallet(ctx, update)
		}

		b.token = nil
	}

	if update.Message.Text == "a" {
		if err := b.db.CreateUser(ctx, update.Message.From.ID); err != nil {
			log.Println(err)
		}
	}

	if update.Message.Text == "g" {
		user, err := b.db.GetUser(ctx, update.Message.From.ID)
		if err != nil {
			log.Println(err)
		}

		log.Println(user.ID, user.UserID, user.Wallets)
	}

	if update.Message.Text == "aw" {
		if err := b.db.AddWallet(ctx, update.Message.From.ID, "test"); err != nil {
			log.Println(err)
		}
	}
}

func (b Bot) handleCommands(ctx context.Context, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		b.processStart(ctx, update)
	case "track":
		b.processTrack(update)
	case "profile":
		b.processProfile(update)
		//default:
		//	msg.Text = "I don't know that command"
	}
}
