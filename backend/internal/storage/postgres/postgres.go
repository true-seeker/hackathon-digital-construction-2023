package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func (s *Storage) GetDb() *sql.DB {
	return s.db
}

func New(storagePath string) (*Storage, error) {
	connStr := "user=postgres password=postgres dbname=hackathon sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE EXTENSION IF NOT EXISTS "pgcrypto";
		
		CREATE TABLE IF NOT EXISTS buildings
		(
			id      uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
			name    varchar(64),
			address varchar(512)
		);
		
		
		
		CREATE TABLE IF NOT EXISTS elevators
		(
			id          uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
			building_id uuid             NOT NULL REFERENCES buildings (id),
			name        varchar(64)      NOT NULL
		);
		
		CREATE TABLE IF NOT EXISTS screens
		(
			id          uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
			elevator_id uuid             NOT NULL REFERENCES elevators (id),
			name        varchar(64)      NOT NULL,
			x           integer          NOT NULL,
			y           integer          NOT NULL
		);
		
		CREATE TABLE IF NOT EXISTS widget_types
		(
			id   uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
			name varchar(64)      NOT NULL
		);
		
		CREATE TABLE IF NOT EXISTS widgets
		(
			id      uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
			name    varchar(64)      NOT NULL,
			type_id uuid             NOT NULL REFERENCES widget_types (id)
		);
		
		CREATE TABLE IF NOT EXISTS screen_widgets
		(
			id           uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
			screen_id    uuid             NOT NULL REFERENCES screens (id),
			widget_id    uuid             NOT NULL REFERENCES widgets (id),
			deleted_date timestamp                 DEFAULT NULL
		);
		`)
	if err != nil {
		panic(err)
	}
	return &Storage{db: db}, nil
}
