package entities

const (
	C = "C"
	F = "F"
)

type Weather struct {
	Temperature *Temperature `json:"temperature,omitempty"`
	Pressure    float32      `json:"pressure,omitempty"`
}

type Temperature struct {
	Value float32 `json:"value,omitempty"`
	Unit  string  `json:"unit,omitempty"`
}
