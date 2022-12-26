package models

type Weather struct {
	CurrentCondition []*struct {
		FeelsLike   string `json:"FeelsLikeC"`
		Cloudcover  string `json:"cloudcover"`
		Humidity    string `json:"humidity"`
		Pressure    string `json:"pressure"`
		Temp        string `json:"temp_C"`
		UvIndex     string `json:"uvIndex"`
		Visibility  string `json:"visibility"`
		WeatherCode string `json:"weatherCode"`
		WeatherDesc []*struct {
			Value string `json:"value"`
		} `json:"weatherDesc"`
		WindDir16Point string `json:"winddir16Point"`
		WindDirDegree  string `json:"winddirDegree"`
		WindSpeedKmPh  string `json:"windspeedKmph"`
	} `json:"current_condition"`
	NearestArea []*struct {
		AreaName []*struct {
			Value string `json:"value"`
		} `json:"areaName"`
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"nearest_area"`
	Weather []*struct {
		Astronomy []*struct {
			Moonrise string `json:"moonrise"`
			Moonset  string `json:"moonset"`
			Sunrise  string `json:"sunrise"`
			Sunset   string `json:"sunset"`
		} `json:"astronomy"`
		Date    string `json:"date"`
		AvgTemp string `json:"avgtempC"`
		MaxTemp string `json:"maxtempC"`
		MinTemp string `json:"mintempC"`
		SunHour string `json:"sunHour"`
	} `json:"weather"`
}
