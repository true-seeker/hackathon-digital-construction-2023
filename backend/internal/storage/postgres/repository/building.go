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

func (b *BuildingRepository) GetAll() ([]*entities.Building, error) {
	rows, err := b.db.Query("select id,name,address,zhk_id from buildings")
	if err != nil {
		// TODO LOGGER
	}
	defer rows.Close()
	var buildings []*entities.Building

	for rows.Next() {
		b := entities.Building{}
		err := rows.Scan(&b.Id, &b.Name, &b.Address, &b.ZhkId)
		if err != nil {
			fmt.Println(err) // TODO LOGGER
			continue
		}
		buildings = append(buildings, &b)
	}
	return buildings, nil
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

func (b *BuildingRepository) GetByZhk(id string) ([]*entities.Building, error) {
	rows, err := b.db.Query("select id,name,address,zhk_id from buildings WHERE zhk_id=$1", id)
	if err != nil {
		// TODO LOGGER
	}
	defer rows.Close()
	var buildings []*entities.Building

	for rows.Next() {
		b := entities.Building{}
		err := rows.Scan(&b.Id, &b.Name, &b.Address, &b.ZhkId)
		if err != nil {
			fmt.Println(err) // TODO LOGGER
			continue
		}
		buildings = append(buildings, &b)
	}
	return buildings, nil
}

func (b *BuildingRepository) GetBuildingByScreenId(screenId string) (*entities.Building, error) {
	row := b.db.QueryRow("select buildings.id,buildings.name,address,zhk_id, latitude, longitude from buildings JOIN elevators e on buildings.id = e.building_id JOIN screens s on e.id = s.elevator_id WHERE s.id=$1", screenId)
	var bd entities.Building
	err := row.Scan(&bd.Id, &bd.Name, &bd.Address, &bd.ZhkId, &bd.Latitude, &bd.Longitude)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}
	return &bd, nil
}
