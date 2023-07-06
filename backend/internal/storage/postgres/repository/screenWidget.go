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
	valueStrings := make([]string, 0, len(*req.ScreenWidgets))
	valueArgs := make([]interface{}, 0, len(*req.ScreenWidgets)*6)
	screenId := ""
	for i, scWidget := range *req.ScreenWidgets {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)", i*6+1, i*6+2, i*6+3, i*6+4, i*6+5, i*6+6))
		valueArgs = append(valueArgs, scWidget.ScreenId)
		valueArgs = append(valueArgs, scWidget.WidgetId)
		valueArgs = append(valueArgs, scWidget.X)
		valueArgs = append(valueArgs, scWidget.Y)
		valueArgs = append(valueArgs, scWidget.XSize)
		valueArgs = append(valueArgs, scWidget.YSize)
		screenId = scWidget.ScreenId
	}
	stmt := fmt.Sprintf("INSERT INTO screen_widgets (screen_id, widget_id, x, y, x_size, y_size)  VALUES %s",
		strings.Join(valueStrings, ","))

	ctx := context.Background()
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}

	_, err = tx.Exec("UPDATE screen_widgets SET deleted_date=now() WHERE screen_id = $1", screenId)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}
	_, err = tx.Exec(stmt, valueArgs...)
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println(err) // TODO LOGGER
	}
	return req.ScreenWidgets, nil
}
