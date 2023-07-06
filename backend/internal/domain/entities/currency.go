package entities

type Currency struct {
	Name  string  `json:"name,omitempty"`
	Value float32 `json:"value,omitempty"`
}
