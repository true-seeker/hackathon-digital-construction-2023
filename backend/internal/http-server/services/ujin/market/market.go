package market

import (
	"backend/internal/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Service struct {
	token string
}

func NewMarketService(token string) *Service {
	return &Service{token: token}
}

type MarketOffers struct {
	Data MarketOfferData `json:"data"`
}
type MarketOfferData struct {
	Offers []Offer `json:"offers"`
}

type Offer struct {
	Id      int        `json:"id"`
	Title   string     `json:"title"`
	Summary string     `json:"summary"`
	Image   OfferImage `json:"image"`
}

type OfferImage struct {
	Url string `json:"url"`
}

type requestBody struct {
	Auth AuthBody `json:"auth"`
}
type AuthBody struct {
	Token string `json:"token"`
	Type  string `json:"type"`
}

func (s *Service) GetMarketOffers() (*MarketOffers, error) {
	client := &http.Client{}
	requestBody := requestBody{Auth: AuthBody{
		Token: s.token,
		Type:  "api-token",
	}}

	jsonBody, _ := json.Marshal(requestBody)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v2/client/resident/marketplace/offers/list", config.Cfg.Ujin.ApiUrl), bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
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

	var offers MarketOffers
	err = json.Unmarshal(body, &offers)
	fmt.Println(1111, string(body))

	if err != nil {
		return nil, err
	}
	return &offers, nil
}
