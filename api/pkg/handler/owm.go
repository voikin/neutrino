package handler

import (
	"net/http"

	owm "github.com/briandowns/openweathermap"
	"github.com/dazai404/neutrino/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) getWeatherByCity(e echo.Context) error {
	input := &models.RequestWithCity{}
	e.Bind(input)
	
	w, err := owm.NewCurrent("C", "ru", h.apiKey)
	if err != nil {
		return nil
	}

	err = w.CurrentByName(input.City)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, *w)
}