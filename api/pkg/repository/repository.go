package repository

import (
	"github.com/voikin/neutrino/models"
	mongoRepository "github.com/voikin/neutrino/pkg/repository/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	TgUserRepository
}

type TgUserRepository interface {
	SaveTgUser(user *models.TgUser) (interface{}, error)
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		TgUserRepository: mongoRepository.NewMongoRepository(db),
	}
}
