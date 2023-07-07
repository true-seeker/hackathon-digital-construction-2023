package datagen

import (
	"backend/internal/domain/entities"
	"backend/internal/http-server/handlers/elevator"
	"backend/internal/http-server/handlers/screen"
	"backend/internal/http-server/services/ujin/complexService"
	"fmt"
	"golang.org/x/exp/rand"
)

type Service struct {
	complexGetter       ComplexGetter
	elevatorsRepository ElevatorsRepository
	screenRepository    ScreenRepository
}

func NewDataGenService(complexGetter ComplexGetter, elevatorsRepository ElevatorsRepository, screenRepository ScreenRepository) *Service {
	return &Service{complexGetter: complexGetter, elevatorsRepository: elevatorsRepository, screenRepository: screenRepository}
}

type ComplexGetter interface {
	GetComplexes() (*complexService.Complex, error)
}

type ElevatorsRepository interface {
	GetByBuilding(buildingId int) ([]*entities.Elevator, error)
	New(request *elevator.SaveRequest) (*entities.Elevator, error)
}

type ScreenRepository interface {
	New(request *screen.SaveRequest) (*entities.Screen, error)
}

func (s *Service) GenerateMockData() {
	cmplx, err := s.complexGetter.GetComplexes()
	if err != nil {
		return
	}

	for _, building := range cmplx.Data.Buildings {
		elevators, err := s.elevatorsRepository.GetByBuilding(building.Id)
		if err != nil {
			continue
		}
		if len(elevators) == 0 {
			e, err := s.elevatorsRepository.New(&elevator.SaveRequest{
				Name:       fmt.Sprintf("Лифт № %d", rand.Intn(10)+1),
				BuildingId: building.Id,
			})
			if err != nil {
				continue
			}
			s.screenRepository.New(&screen.SaveRequest{
				Name:       fmt.Sprintf("Экран № %d", rand.Intn(10)+1),
				ElevatorId: e.Id,
				X:          rand.Intn(250) + 250,
				Y:          rand.Intn(250) + 250,
			})
		}
	}
}
