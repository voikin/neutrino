package models

type TgUser struct {
	FullName string `json:"full_name"`
	Mention string `json:"mention"`
	Locale *struct {
		Language interface{} `json:"language"`
		Territory interface{} `json:"territory"`
		Script interface{} `json:"script"`
		Variant interface{} `json:"variant"`
		Data interface{} `json:"__data"`
		Identifier string `json:"identifier"`
	} `json:"locale"`
}