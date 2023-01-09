package mongo_repo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const tgUsersCollection = "tg_users"

func NewMongoDBConnection(mongoURI string) (*mongo.Database, error) {
	cl, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}
	return cl.Database("neutrino"), nil
}
