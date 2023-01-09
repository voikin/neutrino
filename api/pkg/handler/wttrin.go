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
	if city == "" {
		return echo.ErrBadRequest
	}

	res, err := http.Get(fmt.Sprintf("http://wttr.in/%s?format=j2", city))
	if err != nil {
		return err
	}

	forecast := &models.Forecast{}
	json.NewDecoder(res.Body).Decode(forecast)
	
	curWeather := getCurrentWeather(forecast)

	return e.JSON(http.StatusOK, curWeather)
}

func (h *Handler) getForecastByCity(e echo.Context) error {
	city := e.QueryParam("city")
	if city == "" {
		return echo.ErrBadRequest
	}

	daysString := e.QueryParam("days")
	if daysString == "" {
		return echo.ErrBadRequest
	}

	days, err := strconv.Atoi(daysString)
	if err != nil {
		return err
	}

	if days < 1 || days > 3 {
		return echo.ErrBadRequest
	}

	res, err := http.Get(fmt.Sprintf("http://wttr.in/%s?format=j2", city))
	if err != nil {
		return err
	}

	forecast := &models.Forecast{}
	json.NewDecoder(res.Body).Decode(forecast)

	curForecast := getCurrentForecast(forecast)

	forecastWeather := curForecast.Weather
	curForecast.Weather = nil

	for i := 0; i < days; i++ {
		curForecast.Weather = append(curForecast.Weather, forecastWeather[i])
	}

	return e.JSON(http.StatusOK, curForecast)
}

func getCurrentWeather(forecast *models.Forecast) *models.CurrentWeather {
	weatherDesc := forecast.CurrentCondition[0].WeatherDesc[0]
	forecastCurCond := forecast.CurrentCondition[0]
	curCond := &models.CurrentCondition{
		FeelsLike: forecastCurCond.FeelsLike,
		Cloudcover: forecastCurCond.Cloudcover,
		Humidity: forecastCurCond.Humidity,
		Pressure: forecastCurCond.Pressure,
		Temp: forecastCurCond.Temp,
		UvIndex: forecastCurCond.UvIndex,
		Visibility: forecastCurCond.Visibility,
		WeatherCode: forecastCurCond.WeatherCode,
		WeatherDesc: weatherDesc,
		WindDir16Point: forecastCurCond.WindDir16Point,
		WindDirDegree: forecastCurCond.WindDirDegree,
		WindSpeedKmPh: forecastCurCond.WindSpeedKmPh,
	}
	nearArea := &models.NearestArea{
		AreaName: (*models.AreaName)(forecast.NearestArea[0].AreaName[0]),
		Latitude: forecast.NearestArea[0].Latitude,
		Longitude: forecast.NearestArea[0].Longitude,
	}
	astrForecast := forecast.Weather[0].Astronomy[0]
	astr := &models.Astronomy{
		Moonrise: astrForecast.Moonrise,
		Moonset: astrForecast.Moonset,
		Sunrise: astrForecast.Sunrise,
		Sunset: astrForecast.Sunset,
	}
	weatherForecast := forecast.Weather[0]
	weather := &models.Weather{
		Astronomy: astr,
		Date: weatherForecast.Date,
		AvgTemp: weatherForecast.AvgTemp,
		MaxTemp: weatherForecast.MaxTemp,
		MinTemp: weatherForecast.MinTemp,
		SunHour: weatherForecast.SunHour,
	}
	return &models.CurrentWeather{
		CurrentCondition: curCond,
		NearestArea: nearArea,
		Weather: weather,
	}
}

func getCurrentForecast(forecast *models.Forecast) *models.CurrentForecast {
	weatherDesc := forecast.CurrentCondition[0].WeatherDesc[0]
	forecastCurCond := forecast.CurrentCondition[0]

	curCond := &models.CurrentCondition{
		FeelsLike: forecastCurCond.FeelsLike,
		Cloudcover: forecastCurCond.Cloudcover,
		Humidity: forecastCurCond.Humidity,
		Pressure: forecastCurCond.Pressure,
		Temp: forecastCurCond.Temp,
		UvIndex: forecastCurCond.UvIndex,
		Visibility: forecastCurCond.Visibility,
		WeatherCode: forecastCurCond.WeatherCode,
		WeatherDesc: weatherDesc,
		WindDir16Point: forecastCurCond.WindDir16Point,
		WindDirDegree: forecastCurCond.WindDirDegree,
		WindSpeedKmPh: forecastCurCond.WindSpeedKmPh,
	}

	nearArea := &models.NearestArea{
		AreaName: (*models.AreaName)(forecast.NearestArea[0].AreaName[0]),
		Latitude: forecast.NearestArea[0].Latitude,
		Longitude: forecast.NearestArea[0].Longitude,
	}

	weather := make([]*models.Weather, 0)
	for _, item := range forecast.Weather {
		astr := item.Astronomy[0]
		weather = append(weather, &models.Weather{
			Astronomy: astr,
			Date: item.Date,
			AvgTemp: item.AvgTemp,
			MaxTemp: item.MaxTemp,
			MinTemp: item.MinTemp,
			SunHour: item.SunHour,
		})
	}

	return &models.CurrentForecast{
		CurrentCondition: curCond,
		NearestArea: nearArea,
		Weather: weather,
	}
}