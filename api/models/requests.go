package models

type RequestWithCity struct {
	City string `json:"city"`
}

type RequestWithCityDays struct {
	City string `json:"city"`
	Days int `json:"days"`
}