package entities

type Screen struct {
	Id         string `json:"id,omitempty"`
	ElevatorId string `json:"elevatorId,omitempty"`
	Name       string `json:"name,omitempty"`
	X          int    `json:"x"`
	Y          int    `json:"y"`
}
