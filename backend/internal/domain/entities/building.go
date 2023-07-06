package entities

type Building struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	ZhkId   string `json:"zhk_id,omitempty"`
}
