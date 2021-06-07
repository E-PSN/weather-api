package model

// Weather entity
type Weather struct {
	Temperature string      `json:"temperature"`
	Wind        string      `json:"wind"`
	Description string      `json:"description"`
	Forecast    [4]Forecast `json:"forecast"`
	FeelsLikeC      string `json:"FeelsLikeC"`
	FeelsLikeF      string `json:"FeelsLikeF"`
	Cloudcover      string `json:"cloudcover"`
	Humidity        string `json:"humidity"`
}

// Forecast entity
type Forecast struct {
	Day         string `json:"day"`
	Temperature string `json:"temperature"`
	Wind        string `json:"wind"`
}

// ErrorMessage entity
type ErrorMessage struct {
	Message string `json:"message"`
}
