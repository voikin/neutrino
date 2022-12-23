package main

import (
	"log"
	"os"

	"github.com/dazai404/neutrino/internal/server"
	"github.com/dazai404/neutrino/pkg/handler"
	"github.com/dazai404/neutrino/pkg/logger"
	"github.com/joho/godotenv"
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

	zapConfig := zap.NewDevelopmentConfig()

	logger := logger.NewLogger(zapConfig)

	handler := handler.NewHandler(logger, apiKey)

	srv := new(server.Server)

	panic(srv.Run(port, handler.InitRoutes()))
}
