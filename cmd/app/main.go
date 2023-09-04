package main

import (
	"log"
	"wallet-transaction-notification/internal/app"
	config "wallet-transaction-notification/internal/cfg"
)

func main() {
	cfgPath := "./config/config.yml"
	cfg := config.GetConfig(cfgPath)

	a, err := app.NewApp(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	if err := a.Run(); err != nil {
		log.Fatalln(err)
	}
}
