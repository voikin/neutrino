package main

import (
	"context"
	"fmt"
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
	atlasUser := os.Getenv("ATLAS_USER")
	atlasPwd := os.Getenv("ATLAS_PWD")

	cl, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@neutrino.oicq8th.mongodb.net/?retryWrites=true&w=majority", atlasUser, atlasPwd)))
	if err != nil {
		log.Fatal(err)
	}

	db := cl.Database("neutrino")
	zapConfig := zap.NewDevelopmentConfig()
	repository := repository.NewRepository(db)
	logger := logger.NewLogger(zapConfig)
	handler := handler.NewHandler(repository, logger, apiKey)

	srv := new(server.Server)

	panic(srv.Run(port, handler.InitRoutes()))
}
