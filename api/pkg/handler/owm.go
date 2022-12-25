package handler

import (
	"fmt"
	"net/http"

	owm "github.com/briandowns/openweathermap"
	"github.com/dazai404/neutrino/models"
	"github.com/labstack/echo/v4"
)

func (h *Handler) getWeatherByCity(e echo.Context) error {
	input := &models.RequestWithCity{}

	err := e.Bind(input)
	if err != nil {
		return err
	}

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

func (h *Handler) getForecastByCity(e echo.Context) error {
	input := &models.RequestWithCityDays{}
	err := e.Bind(input)
	if err != nil {
		return err
	}

	w, err := owm.NewForecast("5", "C", "RU", h.apiKey)
	if err != nil {
		return nil
	}

	err = w.DailyByName(input.City, input.Days)
	if err != nil {
		return err
	}

	forecast := w.ForecastWeatherJson.(*owm.Forecast5WeatherData)

	fmt.Println(forecast)

	out := make([]interface{}, 0)

	for i := 0; i < input.Days; i++ {
		out = append(out, forecast.List[i])
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"city": forecast.City,
		"cnt": forecast.Cnt,
		"list": out,
	})
}
