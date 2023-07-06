package news

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Service struct {
	token string
}

func NewNewsService(token string) *Service {
	return &Service{token: token}
}

type News struct {
	Data NewsData
}
type NewsData struct {
	Items []NewsItems
}

type NewsItems struct {
	Title string
	Text  string
	Date  string
}

func (s *Service) GetNews(buildingId int) (*News, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://api-uae-test.ujin.tech/api/v1/news/list?token=%s&d&buildings=%d&type=news", s.token, buildingId), nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var news News
	err = json.Unmarshal(body, &news)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}
	return &news, nil
}
