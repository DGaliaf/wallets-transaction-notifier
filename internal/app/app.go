package app

import (
	"context"
	"log"
	"wallet-transaction-notification/internal/bot"
	"wallet-transaction-notification/internal/cfg"
	"wallet-transaction-notification/internal/database/mongodb"
)

type App struct {
	cfg *cfg.Config
	bot *bot.Bot
	db  *mongodb.Database
}

func NewApp(ctx context.Context, cfg *cfg.Config) (*App, error) {
	log.Println("...start application")

	log.Println("...start database")
	db, err := mongodb.NewDatabase(ctx, cfg)
	if err != nil {
		return nil, err
	}

	tgBot := bot.NewBot(cfg, db)

	return &App{
		cfg: cfg,
		bot: tgBot,
		db:  db,
	}, nil
}

func (a App) Run(ctx context.Context) error {
	log.Println("...start telegram bot")
	if err := a.bot.Run(ctx); err != nil {
		return err
	}

	return nil
}
