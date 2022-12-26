package mongo_repo

import (
	"context"

	"github.com/voikin/neutrino/models"
	"go.mongodb.org/mongo-driver/bson"
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

func (mr *MongoRepository) GetTgUser(id int) (*models.TgUser, error) {
	coll := mr.db.Collection(tgUsersCollection)
	user := &models.TgUser{}

	err := coll.FindOne(
		context.Background(),
		bson.D{{"_id", id}},
	).Decode(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (mr *MongoRepository) UpdateTgUser(user *models.TgUser) error {
	coll := mr.db.Collection(tgUsersCollection)

	_, err := coll.ReplaceOne(
		context.Background(),
		bson.D{{"_id", user.UserId}},
		user,
	)
	if err != nil {
		return err
	}

	return nil
}
