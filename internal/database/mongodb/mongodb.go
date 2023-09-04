package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"wallet-transaction-notification/internal/cfg"
)

type Database struct {
	cfg *cfg.Config
}

func NewDatabase(cfg *cfg.Config) *Database {
	return &Database{cfg: cfg}
}

func (d Database) Run(ctx context.Context) (*mongo.Collection, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/",
		d.cfg.DB.Username, d.cfg.DB.Password, d.cfg.DB.Host, d.cfg.DB.Port)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client.Database("telegram-wallet-notifier").Collection("users"), nil
}
