package app

import (
	"context"
	"golang.org/x/sync/errgroup"
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

func NewApp(cfg *cfg.Config) (*App, error) {
	db := mongodb.NewDatabase(cfg)

	log.Println("...connect to database")
	collection, err := db.Run(context.Background())
	if err != nil {
		return nil, err
	}

	tgBot := bot.NewBot(cfg, collection)

	return &App{
		cfg: cfg,
		bot: tgBot,
		db:  db,
	}, nil
}

func (a App) Run() error {
	log.Println("...start application")

	group, _ := errgroup.WithContext(context.Background())

	group.Go(func() error {
		log.Println("...start telegram bot")
		return a.bot.Run()
	})

	if err := group.Wait(); err != nil {
		return err
	}

	return nil
}
