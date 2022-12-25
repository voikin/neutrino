package handler

import (
	"net/http"

	"github.com/dazai404/neutrino/pkg/logger"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	logger *logger.Logger
	apiKey string
}

func NewHandler(logger *logger.Logger, owmApiKey string) *Handler {
	return &Handler{
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
