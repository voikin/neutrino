package mongo_repo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const tgUsersCollection = "tg_users"

func NewMongoDBConnection(mongoURI string) (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 5)
	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}
	return cl.Database("neutrino"), nil
}
