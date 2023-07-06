package currency

import "backend/internal/domain/entities"

type Service struct {
}

func NewCurrencyService() *Service {
	return &Service{}
}

func (s *Service) GetCurrency() (*[]entities.Currency, error) {
	currencies := []entities.Currency{{
		Name:  "USD",
		Value: 90.35,
	},
		{
			Name:  "EUR",
			Value: 98.114,
		},
		{
			Name:  "BYN",
			Value: 29.83,
		},
		{
			Name:  "KZT",
			Value: 20.123,
		},
	}
	return &currencies, nil
}
