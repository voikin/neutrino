package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/voikin/neutrino/internal/server"
	"github.com/voikin/neutrino/pkg/handler"
	"github.com/voikin/neutrino/pkg/logger"
	"github.com/voikin/neutrino/pkg/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error: no .env file found")
	}
}

func main() {
	apiKey := os.Getenv("OWM_API_KEY")
	port := os.Getenv("PORT")

	cl, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	db := cl.Database("neutrino")
	repository := repository.NewRepository(db)
	zapConfig := zap.NewDevelopmentConfig()
	logger := logger.NewLogger(zapConfig)
	handler := handler.NewHandler(repository, logger, apiKey)

	srv := new(server.Server)

	panic(srv.Run(port, handler.InitRoutes()))
}
