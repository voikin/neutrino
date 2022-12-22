package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dazai404/neutrino/pkg/logger"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("error: no .env file found")
	}
}

func main() {
	apiKey := os.Getenv("OWM_API_KEY")

	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "pong",
		})
	})

	zc := zap.Config{}

	logger := logger.NewLogger(zc)

	e.Use(logger.LoggerMiddleware())

	fmt.Println(apiKey)
	e.Logger.Fatal(e.Start(":8080"))
}
