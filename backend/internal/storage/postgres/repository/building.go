package repository

import (
	"database/sql"
)

type BuildingRepository struct {
	db *sql.DB
}

func NewBuildingRepository(db *sql.DB) *BuildingRepository {
	return &BuildingRepository{db: db}
}
func (b *BuildingRepository) GetBuildingIdByScreenId(screenId string) (int, error) {
	row := b.db.QueryRow("select building_id FROM elevators e JOIN screens s on e.id = s.elevator_id WHERE s.id=$1", screenId)
	var buildingId int
	err := row.Scan(&buildingId)
	if err != nil {
		return 0, err
	}
	return buildingId, nil
}
