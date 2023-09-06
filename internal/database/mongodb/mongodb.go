package mongodb

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"wallet-transaction-notification/internal/cfg"
)

type Database struct {
	cfg        *cfg.Config
	collection *mongo.Collection
}

func NewDatabase(ctx context.Context, cfg *cfg.Config) (*Database, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/",
		cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	collection := client.Database("telegram-wallet-notifier").Collection("users")

	return &Database{
		cfg:        cfg,
		collection: collection,
	}, nil
}

func (d Database) CreateUser(ctx context.Context, userID int64) error {
	user := User{
		UserID:  userID,
		Wallets: make([]string, 0),
	}

	insertedID, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	if insertedID == nil {
		return errors.New("failed to insert user")
	}

	return nil
}

func (d Database) GetUser(ctx context.Context, userID int64) (*GetUser, error) {
	user := &GetUser{}

	if err := d.collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (d Database) AddWallet(ctx context.Context, userID int64, wallet string) error {
	user, err := d.GetUser(ctx, userID)
	if err != nil {
		return err
	}

	user.Wallets = append(user.Wallets, wallet)

	update, err := d.collection.UpdateOne(ctx, bson.M{"user_id": userID}, bson.D{{"$set", user}})
	if err != nil {
		return err
	}

	if update.ModifiedCount == 0 {
		return errors.New(fmt.Sprintf("failed to update user (ID: %d)", userID))
	}

	return nil
}
