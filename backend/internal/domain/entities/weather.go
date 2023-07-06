package entities

import "time"

const (
	C = "C"
	F = "F"
)

const (
	SNOW = "snow"
	RAIN = "rain"
)

type Weather struct {
	Temperature   *Temperature   `json:"temperature,omitempty"`
	Pressure      float32        `json:"pressure,omitempty"`
	Precipitation *Precipitation `json:"precipitation,omitempty"`
}

type Temperature struct {
	Value float32 `json:"value,omitempty"`
	Unit  string  `json:"unit,omitempty"`
}

type Precipitation struct {
	Chance    int       `json:"chance,omitempty"`
	Type      string    `json:"type,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`
}
