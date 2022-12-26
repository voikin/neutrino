package handler

import (
	"fmt"
	"log"
	"net/http"

	owm "github.com/briandowns/openweathermap"
	"github.com/labstack/echo/v4"
	"github.com/voikin/neutrino/models"
)

func (h *Handler) getWeatherByCity(e echo.Context) error {
	input := &models.RequestWithCity{}

	err := e.Bind(input)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(123)

	w, err := owm.NewCurrent("C", "ru", h.apiKey)
	if err != nil {
		log.Println(err)
		return nil
	}
	w.Settings
	fmt.Println(123)

	err = w.CurrentByName(input.City)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(123)

	return e.JSON(http.StatusOK, *w)
}

func (h *Handler) getForecastByCity(e echo.Context) error {
	input := &models.RequestWithCityDays{}
	err := e.Bind(input)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	w, err := owm.NewForecast("5", "C", "RU", h.apiKey)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = w.DailyByName(input.City, input.Days)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	forecast := w.ForecastWeatherJson.(*owm.Forecast5WeatherData)

	out := make([]interface{}, 0)

	for i := 0; i < input.Days; i++ {
		out = append(out, forecast.List[i])
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"city": forecast.City,
		"cnt":  forecast.Cnt,
		"list": out,
	})
}

func owmForecastPing() error {
	_, err := http.Get("https://api.openweathermap.org/data/2.5/forecast?appid=91234a30076102d81d6c6b6a6f362b2c&mode=json&units=C&lang=RU&cnt=3")
	if err != nil {
		return err
	}
	return nil
}
