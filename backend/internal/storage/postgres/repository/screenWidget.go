package repository

import (
	"backend/internal/domain/entities"
	"backend/internal/http-server/handlers/screenWidget"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type ScreenWidgetRepository struct {
	db *sql.DB
}

func NewScreenWidgetRepository(db *sql.DB) *ScreenWidgetRepository {
	return &ScreenWidgetRepository{db: db}
}

func (s *ScreenWidgetRepository) Save(req *screenWidget.SaveRequest) (*[]entities.ScreenWidget, error) {
	fieldCount := 10
	valueStrings := make([]string, 0, len(*req.ScreenWidgets))
	valueArgs := make([]interface{}, 0, len(*req.ScreenWidgets)*fieldCount)
	for i, scWidget := range *req.ScreenWidgets {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)",
			i*fieldCount+1, i*fieldCount+2, i*fieldCount+3, i*fieldCount+4, i*fieldCount+5, i*fieldCount+6, i*fieldCount+7, i*fieldCount+8, i*fieldCount+9, i*fieldCount+10))
		valueArgs = append(valueArgs, req.ScreenId)
		valueArgs = append(valueArgs, scWidget.I)
		valueArgs = append(valueArgs, scWidget.X)
		valueArgs = append(valueArgs, scWidget.Y)
		valueArgs = append(valueArgs, scWidget.W)
		valueArgs = append(valueArgs, scWidget.H)
		valueArgs = append(valueArgs, scWidget.MinW)
		valueArgs = append(valueArgs, scWidget.MinH)
		valueArgs = append(valueArgs, scWidget.Moved)
		valueArgs = append(valueArgs, scWidget.Static)
	}
	stmt := fmt.Sprintf("INSERT INTO screen_widgets (screen_id, i, x, y, w, h, min_w, min_h, moved, static)  VALUES %s",
		strings.Join(valueStrings, ","))

	ctx := context.Background()
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec("UPDATE screen_widgets SET deleted_date=now() WHERE screen_id = $1 AND deleted_date is null", req.ScreenId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	_, err = tx.Exec(stmt, valueArgs...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return req.ScreenWidgets, nil
}

func (s *ScreenWidgetRepository) Get(screenId string) ([]*entities.ScreenWidget, error) {
	rows, err := s.db.Query("select id, screen_id, i, x, y, w, h, min_w, min_h, moved, static from screen_widgets "+
		"WHERE screen_id=$1 AND deleted_date IS NULL", screenId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var screenWidgets []*entities.ScreenWidget

	for rows.Next() {
		b := entities.ScreenWidget{}
		err := rows.Scan(&b.Id, &b.ScreenId, &b.I, &b.X, &b.Y, &b.W, &b.H, &b.MinW, &b.MinH, &b.Moved, &b.Static)
		if err != nil {
			continue
		}
		screenWidgets = append(screenWidgets, &b)
	}
	return screenWidgets, nil
}
