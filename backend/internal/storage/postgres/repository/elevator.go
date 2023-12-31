package repository

import (
	"backend/internal/domain/entities"
	"backend/internal/http-server/handlers/elevator"
	"database/sql"
)

type ElevatorRepository struct {
	db *sql.DB
}

func NewElevatorRepository(db *sql.DB) *ElevatorRepository {
	return &ElevatorRepository{db: db}
}

func (b *ElevatorRepository) GetAll() ([]*entities.Elevator, error) {
	rows, err := b.db.Query("select id,name,building_id from elevators")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var elevators []*entities.Elevator

	for rows.Next() {
		e := entities.Elevator{}
		err := rows.Scan(&e.Id, &e.Name, &e.BuildingId)
		if err != nil {
			continue
		}
		elevators = append(elevators, &e)
	}
	return elevators, nil
}

func (b *ElevatorRepository) Get(id string) (*entities.Elevator, error) {
	row := b.db.QueryRow("select id, name, building_id from elevators WHERE id=$1", id)
	var e entities.Elevator
	err := row.Scan(&e.Id, &e.Name, &e.BuildingId)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (b *ElevatorRepository) New(request *elevator.SaveRequest) (*entities.Elevator, error) {
	id := ""
	_ = b.db.QueryRow("insert into elevators(name, building_id) values ($1, $2) RETURNING id",
		request.Name, request.BuildingId).Scan(&id)

	e, err := b.Get(id)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (b *ElevatorRepository) Update(request *elevator.UpdateRequest) (*entities.Elevator, error) {
	_, err := b.db.Exec("update elevators SET name=$2, building_id=$3 WHERE id=$1",
		request.Id, request.Name, request.BuildingId)
	if err != nil {
		return nil, err
	}
	bd, err := b.Get(request.Id)
	if err != nil {
		return nil, err
	}

	return bd, nil
}

func (b *ElevatorRepository) GetByBuilding(buildingId int) ([]*entities.Elevator, error) {
	rows, err := b.db.Query("select id,name,building_id from elevators WHERE building_id=$1", buildingId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var elevators []*entities.Elevator

	for rows.Next() {
		e := entities.Elevator{}
		err := rows.Scan(&e.Id, &e.Name, &e.BuildingId)
		if err != nil {
			continue
		}
		elevators = append(elevators, &e)
	}
	return elevators, nil
}
