package repository

import (
	"backend/internal/domain/entities"
	"backend/internal/http-server/handlers/building"
	"database/sql"
	"fmt"
)

type BuildingRepository struct {
	db *sql.DB
}

func NewBuildingRepository(db *sql.DB) *BuildingRepository {
	return &BuildingRepository{db: db}
}

func (b *BuildingRepository) Get(id string) (*entities.Building, error) {
	row := b.db.QueryRow("select id, name, address, zhk_id,latitude, longitude from buildings WHERE id=$1", id)
	var bd entities.Building
	err := row.Scan(&bd.Id, &bd.Name, &bd.Address, &bd.ZhkId, &bd.Latitude, &bd.Longitude)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}

	return &bd, nil
}

func (b *BuildingRepository) New(request *building.SaveRequest) (*entities.Building, error) {
	id := ""
	_ = b.db.QueryRow("insert into buildings (name, address, zhk_id, latitude, longitude) values ($1, $2, $3, $4, $5) RETURNING id",
		request.Name, request.Address, request.ZhkId, request.Latitude, request.Longitude).Scan(&id)

	bd, err := b.Get(id)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}

	return bd, nil
}

func (b *BuildingRepository) Update(request *building.UpdateRequest) (*entities.Building, error) {
	_, err := b.db.Exec("update buildings SET name=$2, address=$3, zhk_id=$4 WHERE id=$1",
		request.Id, request.Name, request.Address, request.ZhkId)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}
	bd, err := b.Get(request.Id)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}

	return bd, nil
}

func (b *BuildingRepository) GetBuildingIdByScreenId(screenId string) (int, error) {
	row := b.db.QueryRow("select building_id FROM elevators e JOIN screens s on e.id = s.elevator_id WHERE s.id=$1", screenId)
	var buildingId int
	err := row.Scan(&buildingId)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}
	return buildingId, nil
}
