package weather

import "backend/internal/domain/entities"

type Service struct {
}

func NewWeatherService() *Service {
	return &Service{}
}

func (s *Service) GetWeather() (*entities.Weather, error) {
	return &entities.Weather{
		Temperature: &entities.Temperature{
			Value: 52,
			Unit:  entities.C,
		},
		Pressure: 52,
	}, nil
}
