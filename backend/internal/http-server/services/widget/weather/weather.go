package weather

import (
	"backend/internal/domain/entities"
	"time"
)

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
		Precipitation: &entities.Precipitation{
			Chance:    25,
			Type:      entities.RAIN,
			StartTime: time.Now(),
			EndTime:   time.Now().Add(time.Hour * 3),
		},
	}, nil
}
