package entities

import "time"

type WidgetType struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Widget struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	TypeId string `json:"type_id,omitempty"`
}

type ScreenWidget struct {
	Id          string    `json:"id"`
	ScreenId    string    `json:"screen_id"`
	I           string    `json:"i"`
	X           int       `json:"x"`
	Y           int       `json:"y"`
	W           int       `json:"w"`
	H           int       `json:"h"`
	MinW        int       `json:"min_w"`
	MinH        int       `json:"min_h"`
	Moved       bool      `json:"moved"`
	Static      bool      `json:"static"`
	DeletedDate time.Time `json:"deleted_date"`
}
