package repository

import (
	"backend/internal/domain/entities"
	"backend/internal/http-server/handlers/screen"
	"database/sql"
	"fmt"
)

type ScreenRepository struct {
	db *sql.DB
}

func NewScreenRepository(db *sql.DB) *ScreenRepository {
	return &ScreenRepository{db: db}
}

func (b *ScreenRepository) GetAll() ([]*entities.Screen, error) {
	rows, err := b.db.Query("select id,name,elevator_id,x,y from screen")
	if err != nil {
		// TODO LOGGER
	}
	defer rows.Close()
	var screens []*entities.Screen

	for rows.Next() {
		e := entities.Screen{}
		err := rows.Scan(&e.Id, &e.Name, &e.ElevatorId, &e.X, &e.Y)
		if err != nil {
			fmt.Println(err) // TODO LOGGER
			continue
		}
		screens = append(screens, &e)
	}
	return screens, nil
}

func (b *ScreenRepository) Get(id string) (*entities.Screen, error) {
	row := b.db.QueryRow("select id, name, elevator_id,x,y from screen WHERE id=$1", id)
	var e entities.Screen
	err := row.Scan(&e.Id, &e.Name, &e.ElevatorId, &e.X, &e.Y)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}

	return &e, nil
}

func (b *ScreenRepository) New(request *screen.SaveRequest) (*entities.Screen, error) {
	id := ""
	_ = b.db.QueryRow("insert into screen(name, elevator_id,x,y) values ($1, $2, $3,$4) RETURNING id",
		request.Name, request.ElevatorId, request.X, request.Y).Scan(&id)

	e, err := b.Get(id)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}

	return e, nil
}

func (b *ScreenRepository) Update(request *screen.UpdateRequest) (*entities.Screen, error) {
	_, err := b.db.Exec("update screen SET name=$2, elevator_id=$3, x=$4, y=$5 WHERE id=$1",
		request.Id, request.Name, request.ElevatorId, request.X, request.Y,
	)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}
	bd, err := b.Get(request.Id)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}

	return bd, nil
}

func (b *ScreenRepository) GetByElevator(elevatorId string) ([]*entities.Screen, error) {
	rows, err := b.db.Query("select id,name,elevator_id,x,y from screen WHERE elevator_id=$1", elevatorId)
	if err != nil {
		// TODO LOGGER
	}
	defer rows.Close()
	var screens []*entities.Screen

	for rows.Next() {
		e := entities.Screen{}
		err := rows.Scan(&e.Id, &e.Name, &e.ElevatorId, &e.X, &e.Y)
		if err != nil {
			fmt.Println(err) // TODO LOGGER
			continue
		}
		screens = append(screens, &e)
	}
	return screens, nil
}
