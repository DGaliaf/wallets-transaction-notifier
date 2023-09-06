package mongodb

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserID  int64    `bson:"user_id"`
	Wallets []string `bson:"wallets"`
}

type GetUser struct {
	ID      primitive.ObjectID `bson:"_id"`
	UserID  int64              `bson:"user_id"`
	Wallets []string           `bson:"wallets"`
}
