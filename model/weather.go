package model

// Weather entity
type Weather struct {
	FeelsLikeC      string `json:"FeelsLikeC"`
	FeelsLikeF      string `json:"FeelsLikeF"`
	Cloudcover      string `json:"cloudcover"`
	Humidity        string `json:"humidity"`
	ObservationTime string `json:"observation_time"`
	PrecipMM        string `json:"precipMM"`
	Pressure        string `json:"pressure"`
	TempC           string `json:"temp_C"`
	TempF           string `json:"temp_F"`
	UvIndex         int    `json:"uvIndex"`
	Visibility      string `json:"visibility"`
	WeatherCode     string `json:"weatherCode"`	
	Temperature string      `json:"temperature"`
	Wind        string      `json:"wind"`
	Description string      `json:"description"`
	Forecast    [3]Forecast `json:"forecast"`
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
