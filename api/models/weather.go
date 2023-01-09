package models

type Forecast struct {
	CurrentCondition []*struct {
		FeelsLike   string `json:"feelsLikeC"`
		Cloudcover  string `json:"cloudcover"`
		Humidity    string `json:"humidity"`
		Pressure    string `json:"pressure"`
		Temp        string `json:"temp_C"`
		UvIndex     string `json:"uvIndex"`
		Visibility  string `json:"visibility"`
		WeatherCode string `json:"weatherCode"`
		WeatherDesc []*WeatherDesc `json:"weatherDesc"`
		WindDir16Point string `json:"winddir16Point"`
		WindDirDegree  string `json:"winddirDegree"`
		WindSpeedKmPh  string `json:"windspeedKmph"`
	} `json:"current_condition"`
	NearestArea []*struct {
		AreaName []*AreaName `json:"areaName"`
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"nearest_area"`
	Weather []*struct {
		Astronomy []*Astronomy `json:"astronomy"`
		Date    string `json:"date"`
		AvgTemp string `json:"avgtempC"`
		MaxTemp string `json:"maxtempC"`
		MinTemp string `json:"mintempC"`
		SunHour string `json:"sunHour"`
	} `json:"weather"`
}

type CurrentWeather struct {
	CurrentCondition *CurrentCondition `json:"current_condition"`
	NearestArea *NearestArea `json:"nearest_area"`
	Weather *Weather `json:"weather"`
}

type CurrentForecast struct {
	CurrentCondition *CurrentCondition `json:"current_condition"`
	NearestArea *NearestArea `json:"nearest_area"`
	Weather []*Weather `json:"weather"`
}

type WeatherDesc struct {
	Value string `json:"value"`
}

type CurrentCondition struct {
	FeelsLike      string         `json:"feelsLikeC"`
	Cloudcover     string         `json:"cloudcover"`
	Humidity       string         `json:"humidity"`
	Pressure       string         `json:"pressure"`
	Temp           string         `json:"temp_C"`
	UvIndex        string         `json:"uvIndex"`
	Visibility     string         `json:"visibility"`
	WeatherCode    string         `json:"weatherCode"`
	WeatherDesc    *WeatherDesc `json:"weatherDesc"`
	WindDir16Point string         `json:"winddir16Point"`
	WindDirDegree  string         `json:"winddirDegree"`
	WindSpeedKmPh  string         `json:"windspeedKmph"`
}

type AreaName struct {
	Value string `json:"value"`
}

type NearestArea struct {
	AreaName *AreaName `json:"areaName"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Astronomy struct {
	Moonrise string `json:"moonrise"`
	Moonset  string `json:"moonset"`
	Sunrise  string `json:"sunrise"`
	Sunset   string `json:"sunset"`
}

type Weather struct {
	Astronomy *Astronomy `json:"astronomy"`
	Date    string `json:"date"`
	AvgTemp string `json:"avgtempC"`
	MaxTemp string `json:"maxtempC"`
	MinTemp string `json:"mintempC"`
	SunHour string `json:"sunHour"`
}