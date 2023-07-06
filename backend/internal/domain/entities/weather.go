package entities

const (
	C = "C"
	F = "F"
)

const (
	SNOW = "snow"
	RAIN = "rain"
)

type Weather struct {
	TemperatureNow Temperature `json:"temperature_now"`
	Pressure       int         `json:"pressure"`
	Condition      string      `json:"condition"`
	FeelsLike      int         `json:"feels_like"`
	Forecast       []Temperature
}

type Temperature struct {
	Value     int    `json:"value"`
	Date      string `json:"date"`
	Condition string `json:"condition"`
}
