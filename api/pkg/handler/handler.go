package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/voikin/neutrino/pkg/logger"
	"github.com/voikin/neutrino/pkg/repository"
)

type Handler struct {
	repo   *repository.Repository
	logger *logger.Logger
	apiKey string
}

func NewHandler(repo *repository.Repository, logger *logger.Logger, owmApiKey string) *Handler {
	return &Handler{
		repo:   repo,
		logger: logger,
		apiKey: owmApiKey,
	}
}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()
	e.Use(h.logger.LoggerMiddleware())

	e.GET("/ping", h.ping)

	api := e.Group("/api")
	{
		api.GET("/weather-by-city", h.getWeatherByCity)
		api.GET("/forecast-by-city", h.getForecastByCity)
		users := api.Group("/users")
		{
			users.POST("/save", h.saveTgUser)
		}
	}

	return e
}

func (h *Handler) ping(c echo.Context) error {
	return c.JSON(http.StatusOK, jsonMap{
		"message": "pong",
	})
}
