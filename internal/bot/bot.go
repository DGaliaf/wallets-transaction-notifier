package bot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"wallet-transaction-notification/internal/cfg"
	"wallet-transaction-notification/internal/database/mongodb"
)

type Bot struct {
	cfg *cfg.Config
	db  *mongodb.Database
	bot *tgbotapi.BotAPI

	token *string
}

func NewBot(cfg *cfg.Config, db *mongodb.Database) *Bot {
	return &Bot{
		cfg:   cfg,
		db:    db,
		token: new(string),
	}
}

func (b Bot) Run(ctx context.Context) error {
	var err error
	if b.bot, err = tgbotapi.NewBotAPI(b.cfg.Bot.Token); err != nil {
		return err
	}

	b.bot.Debug = true

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	b.handleUpdates(ctx, u)

	return nil
}
