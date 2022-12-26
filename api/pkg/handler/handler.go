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
		api.POST("/weather-by-city", h.getWeatherByCity)
		api.POST("/forecast-by-city", h.getForecastByCity)
	}

	return e
}

func (h *Handler) ping(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "pong",
	})
}
