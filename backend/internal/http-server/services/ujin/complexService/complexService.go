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

func (s *Service) GetComplexes() (*Complex, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api-uae-test.ujin.tech/api/v1/buildings/get-list-crm?token=%s", s.token), nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var complex Complex
	err = json.Unmarshal(body, &complex)
	if err != nil {
		return nil, err
	}
	return &complex, nil
}

func (s *Service) GetBuildingById(buildingId int) (*BuildingInfo, error) {
	complex, err := s.GetComplexes()
	if err != nil {
		return nil, err
	}
	for _, building := range complex.Data.Buildings {
		if building.Id == buildingId {
			return &building.BuildingInfo, nil
		}
	}
	return nil, err
}
