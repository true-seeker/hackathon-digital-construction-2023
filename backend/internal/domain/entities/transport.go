package entities

const (
	BUS    = "BUS"
	TRAM   = "TRAM"
	SUBWAY = "SUBWAY"
)

type Transport struct {
	Station  string     `json:"station,omitempty"`
	Type     string     `json:"type,omitempty"`
	Arrivals *[]Arrival `json:"arrivals,omitempty"`
}

type Arrival struct {
	Number   int `json:"number,omitempty"`
	TimeLeft int `json:"time_left,omitempty"`
}
