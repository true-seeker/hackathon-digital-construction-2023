package complexService

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Service struct {
	token string
}

func NewComplexService(token string) *Service {
	return &Service{token: token}
}

type Complex struct {
	Data ComplexData `json:"data"`
}
type ComplexData struct {
	Buildings []Building `json:"buildings,omitempty"`
}

type Building struct {
	Id           int          `json:"id"`
	Complex      ComplexInfo  `json:"complex"`
	BuildingInfo BuildingInfo `json:"building"`
}

type ComplexInfo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type BuildingInfo struct {
	Id      int             `json:"id"`
	Title   string          `json:"title"`
	Address BuildingAddress `json:"address"`
}

type BuildingAddress struct {
	FullAddress string `json:"fullAddress"`
}

func (s *Service) GetComplexes() (Complex, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://api-uae-test.ujin.tech/api/v1/buildings/get-list-crm?token=%s", s.token), nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var complex Complex
	err = json.Unmarshal(body, &complex)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}

	fmt.Println(complex)
	return complex, nil
}
