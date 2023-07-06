package repository

import (
	"backend/internal/domain/entities"
	"backend/internal/http-server/handlers/zhk"
	"database/sql"
	"fmt"
)

type ZhkRepository struct {
	db *sql.DB
}

func NewZhkRepository(db *sql.DB) *ZhkRepository {
	return &ZhkRepository{db: db}
}

func (b *ZhkRepository) Get(id string) (*entities.Zhk, error) {
	row := b.db.QueryRow("select id, name from zhks WHERE id=$1", id)
	var bd entities.Zhk
	err := row.Scan(&bd.Id, &bd.Name)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}

	return &bd, nil
}

func (b *ZhkRepository) New(request *zhk.SaveRequest) (*entities.Zhk, error) {
	id := ""
	_ = b.db.QueryRow("insert into zhks (name) values ($1) RETURNING id",
		request.Name).Scan(&id)

	bd, err := b.Get(id)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}

	return bd, nil
}

func (b *ZhkRepository) Update(request *zhk.UpdateRequest) (*entities.Zhk, error) {
	_, err := b.db.Exec("update zhks SET name=$2 WHERE id=$1",
		request.Id, request.Name)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}
	bd, err := b.Get(request.Id)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}

	return bd, nil
}
