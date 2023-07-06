package weather

import (
	"backend/internal/domain/entities"
)

type ServiceMock struct {
}

func NewWeatherMockService() *ServiceMock {
	return &ServiceMock{}
}

func (s *ServiceMock) GetWeather() (*entities.Weather, error) {
	return nil, nil
}
