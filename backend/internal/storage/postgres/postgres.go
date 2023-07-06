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

	_, err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto";
							
							CREATE TABLE IF NOT EXISTS zhks
							(
								id   uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
								name varchar(64)
							);
							
							CREATE TABLE IF NOT EXISTS buildings
							(
								id      uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
								name    varchar(64),
								address varchar(512),
								zhk_id  uuid             NOT NULL REFERENCES zhks (id)
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
								id   uuid PRIMARY KEY 		NOT NULL DEFAULT gen_random_uuid(),
								name varchar(64) UNIQUE    NOT NULL
							);
							
							CREATE TABLE IF NOT EXISTS widgets
							(
								id      uuid PRIMARY KEY 	NOT NULL DEFAULT gen_random_uuid(),
								name    varchar(64) UNIQUE	NOT NULL,
								type_id uuid             	NOT NULL REFERENCES widget_types (id)
							);
							
							CREATE TABLE IF NOT EXISTS screen_widgets
							(
								id           uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
								screen_id    uuid             NOT NULL REFERENCES screens (id),
								widget_id    uuid             NOT NULL REFERENCES widgets (id),
								x            int              NOT NULL,
    							y            int              NOT NULL,
								x_size       int              NOT NULL,
								y_size       int              NOT NULL,
								deleted_date timestamp                 DEFAULT NULL
							);
							INSERT INTO widget_types(name)
							VALUES ('Погода'),
								   ('Курсы валют'),
								   ('Реклама'),
								   ('Транспорт'),
								   ('Новости')
							ON CONFLICT DO NOTHING;

							INSERT INTO widgets (name, type_id)
							VALUES ('Погода', '161a7cb9-80b7-4504-8d9c-f08c828e41a4'),
								   ('Курсы валют', '6846f77a-f030-45eb-a9f3-cadec01c4d51'),
								   ('Реклама', '18b8b45d-82fd-4733-b06e-106649c266d8'),
								   ('Транспорт', '8948742d-3104-4711-9d00-749867407b34'),
								   ('Новости', '90caa3a5-3018-4f69-b03b-c5c313d93ed2') ON CONFLICT DO NOTHING;
		`)
	if err != nil {
		panic(err)
	}
	return &Storage{db: db}, nil
}
