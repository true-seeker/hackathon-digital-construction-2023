package news

import (
	"backend/internal/config"
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
	Data NewsData `json:"data"`
}
type NewsData struct {
	Items []NewsItems `json:"items"`
}

type NewsItems struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Date  string `json:"date"`
}

func (s *Service) GetNews(buildingId int) (*News, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/news/list?token=%s&d&buildings=%d&type=news", config.Cfg.Ujin.ApiUrl, s.token, buildingId), nil)
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

	var news News
	err = json.Unmarshal(body, &news)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}
	return &news, nil
}
