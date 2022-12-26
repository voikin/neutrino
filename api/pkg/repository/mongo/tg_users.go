package mongo_repo

import (
	"context"

	"github.com/voikin/neutrino/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	db *mongo.Database
}

func NewMongoRepository(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		db: db,
	}
}

func (mr *MongoRepository) SaveTgUser(user *models.TgUser) (interface{}, error) {
	collection := mr.db.Collection(tgUsersCollection)
	res, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}
