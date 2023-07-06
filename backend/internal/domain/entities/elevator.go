package entities

type Elevator struct {
	Id         string `json:"id,omitempty"`
	BuildingId string `json:"buildingId,omitempty"`
	Name       string `json:"name,omitempty"`
}
