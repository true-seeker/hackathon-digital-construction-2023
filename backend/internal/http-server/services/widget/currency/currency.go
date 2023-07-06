package currency

import (
	"backend/internal/domain/entities"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Service struct {
}

func NewCurrencyService() *Service {
	return &Service{}
}

type Currency struct {
	NewAmount float32 `json:"new_amount,omitempty"`
}

func (s *Service) GetCurrency() (*[]entities.Currency, error) {
	currencySymbols := []string{"USD", "EUR"}
	var currencies []entities.Currency
	for _, currency := range currencySymbols {
		client := &http.Client{}
		req, _ := http.NewRequest("GET", fmt.Sprintf("https://api.api-ninjas.com/v1/convertcurrency?have=%s&want=RUB&amount=1", currency), nil)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err) // TODO LOGGER
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)

		var cur Currency
		err = json.Unmarshal(body, &cur)
		if err != nil {
			fmt.Println(err) // TODO LOGGER
		}

		currencies = append(currencies, entities.Currency{
			Name:  currency,
			Value: cur.NewAmount,
		})
	}
	return &currencies, nil
}
