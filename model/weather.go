package model

// Weather entity
type Weather struct {
	FeelsLikeC      string `json:"FeelsLikeC"`
	FeelsLikeF      string `json:"FeelsLikeF"`
	Cloudcover      string `json:"cloudcover"`
	Humidity        string `json:"humidity"`
	Temperature 	string      `json:"temperature"`
	Wind        	string      `json:"wind"`
	Description 	string      `json:"description"`
	Forecast    	[3]Forecast `json:"forecast"`
	
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
