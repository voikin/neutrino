package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/voikin/neutrino/models"
)

func (h *Handler) getWeatherByCity(e echo.Context) error {
	city := e.QueryParam("city")
	fmt.Println("city: ", city)

	res, err := http.Get(fmt.Sprintf("http://wttr.in/%s?format=j2", city))
	if err != nil {
		return err
	}

	curWeather := &models.Weather{}
	json.NewDecoder(res.Body).Decode(curWeather)

	cur := curWeather.Weather[0]
	curWeather.Weather = nil
	curWeather.Weather = append(curWeather.Weather, cur)

	return e.JSON(http.StatusOK, curWeather)
}

func (h *Handler) getForecastByCity(e echo.Context) error {
	city := e.QueryParam("city")
	daysString := e.QueryParam("days")
	days, err := strconv.Atoi(daysString)
	if err != nil {
		return err
	}

	if days < 1 && days > 3 {
		return echo.ErrBadRequest
	}

	res, err := http.Get(fmt.Sprintf("http://wttr.in/%s?format=j2", city))
	if err != nil {
		return err
	}

	curWeather := &models.Weather{}
	json.NewDecoder(res.Body).Decode(curWeather)

	forecast := curWeather.Weather
	curWeather.Weather = nil

	for i := 0; i < days; i++ {
		curWeather.Weather = append(curWeather.Weather, forecast[i])
	}

	return e.JSON(http.StatusOK, curWeather)
}
