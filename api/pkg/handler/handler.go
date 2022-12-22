package handler

import (
	"net/http"

	"github.com/dazai404/neutrino/pkg/logger"
	"github.com/labstack/echo/v4"
)

type Handler struct{
	Logger *logger.Logger
}

func NewHandler(logger *logger.Logger) *Handler {
	return &Handler{
		Logger: logger,
	}
}

func (h *Handler) InitRoutes() *echo.Echo {
	e := echo.New()
	e.Use(h.Logger.LoggerMiddleware())
	
	e.GET("/ping", h.ping)

	return e
}

func (h *Handler) ping(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "pong",
	})
}