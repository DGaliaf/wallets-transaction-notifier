package main

import (
	"context"
	"log"
	"wallet-transaction-notification/internal/app"
	config "wallet-transaction-notification/internal/cfg"
)

// TODO: Check for duplicate users
// 		 Make wallet validation

func main() {
	cfgPath := "./config/config.yml"
	cfg := config.GetConfig(cfgPath)

	ctx := context.Background()

	a, err := app.NewApp(ctx, cfg)
	if err != nil {
		log.Fatalln(err)
	}

	if err := a.Run(ctx); err != nil {
		log.Fatalln(err)
	}
}
