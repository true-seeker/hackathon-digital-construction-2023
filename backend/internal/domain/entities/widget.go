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
	Id          string    `json:"id,omitempty"`
	ScreenId    string    `json:"screen_id,omitempty"`
	WidgetId    string    `json:"widget_id,omitempty"`
	X           int       `json:"x,omitempty"`
	Y           int       `json:"y,omitempty"`
	XSize       int       `json:"x_size,omitempty"`
	YSize       int       `json:"y_size,omitempty"`
	DeletedDate time.Time `json:"deleted_date"`
}
