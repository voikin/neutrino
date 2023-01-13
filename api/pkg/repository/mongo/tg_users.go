package mongo_repo

import (
	"context"
	"time"

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

func (mr *MongoRepository) SaveTgUser(user *models.TgUser) error {
	collection := mr.db.Collection(tgUsersCollection)
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func (mr *MongoRepository) UpdateTgUser(user *models.TgUser) error {
	coll := mr.db.Collection(tgUsersCollection)

	_, err := coll.ReplaceOne(
		context.Background(),
		bson.D{{Key: "_id", Value: user.UserId}},
		user,
	)
	if err != nil {
		return err
	}

	return nil
}

func (mr *MongoRepository) GetTgUser(id int) (*models.TgUser, error) {
	coll := mr.db.Collection(tgUsersCollection)
	filter := bson.D{{Key: "_id", Value: id}}
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 10)
	user := &models.TgUser{}
	print(id)
	err := coll.FindOne(ctx, filter, nil).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (mr *MongoRepository) DeleteTgUser(id int) error {
	coll := mr.db.Collection(tgUsersCollection)
	filter := bson.D{{Key: "_id", Value: id}}
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 5)

	_, err := coll.DeleteOne(ctx, filter, nil)
	if err != nil {
		return err
	}

	return nil
}
