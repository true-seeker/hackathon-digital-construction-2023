package entities

type Building struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	ZhkId     int     `json:"zhk_id"`
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}
