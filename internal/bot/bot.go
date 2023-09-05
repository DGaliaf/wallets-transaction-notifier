package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"wallet-transaction-notification/internal/cfg"
)

type Bot struct {
	cfg        *cfg.Config
	collection *mongo.Collection
	bot        *tgbotapi.BotAPI
}

func NewBot(cfg *cfg.Config, collection *mongo.Collection) *Bot {
	return &Bot{
		cfg:        cfg,
		collection: collection,
	}
}

func (b Bot) Run() error {
	var err error
	if b.bot, err = tgbotapi.NewBotAPI(b.cfg.Bot.Token); err != nil {
		return err
	}

	b.bot.Debug = true

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	b.handleUpdates(u)

	return nil
}
