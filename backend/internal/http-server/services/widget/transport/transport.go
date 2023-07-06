package transport

import "backend/internal/domain/entities"

type Service struct {
}

func NewTransportService() *Service {
	return &Service{}
}

func (s *Service) GetTransport() (*[]entities.Transport, error) {
	transport := []entities.Transport{{
		Station: "ул. Пушкина",
		Type:    entities.BUS,
		Arrivals: &[]entities.Arrival{{
			Number:   23,
			TimeLeft: 145,
		},
			{
				Number:   2,
				TimeLeft: 23,
			},
			{
				Number:   5,
				TimeLeft: 199,
			},
		},
	},
		{
			Station: "ул. Ленина",
			Type:    entities.SUBWAY,
			Arrivals: &[]entities.Arrival{{
				Number:   1,
				TimeLeft: 26,
			},
				{
					Number:   2,
					TimeLeft: 40,
				},
			},
		},
		{
			Station: "ул. Екатерининская",
			Type:    entities.TRAM,
			Arrivals: &[]entities.Arrival{{
				Number:   5,
				TimeLeft: 60,
			},
			},
		},
	}
	return &transport, nil
}
