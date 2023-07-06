package repository

import (
	"backend/internal/domain/entities"
	"database/sql"
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
		return nil, err
	}

	return &bd, nil
}
