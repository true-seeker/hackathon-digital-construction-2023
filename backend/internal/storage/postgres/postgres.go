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

func New() (*Storage, error) {
	connStr := "user=postgres password=postgres dbname=hackathon sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto";
							
							CREATE TABLE IF NOT EXISTS elevators
							(
								id          uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
								building_id integer          NOT NULL ,
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
							
							
							CREATE TABLE IF NOT EXISTS widgets
							(
								id      uuid PRIMARY KEY 	NOT NULL DEFAULT gen_random_uuid(),
								name    varchar(64) UNIQUE	NOT NULL
							);
							
							CREATE TABLE IF NOT EXISTS screen_widgets
							(
								id           uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
								screen_id    uuid             NOT NULL REFERENCES screens (id),
								i            uuid             NOT NULL REFERENCES widgets (id),
								x            int              NOT NULL,
								y            int              NOT NULL,
								w            int              NOT NULL,
								h            int              NOT NULL,
								min_w        int              NOT NULL,
								min_h        int              NOT NULL,
								moved        bool                      DEFAULT FALSE,
								static       bool                      DEFAULT FALSE,
								deleted_date timestamp                 DEFAULT NULL
							);
							INSERT INTO widgets (id, name)
							VALUES ('d30bb91e-6718-4380-a196-9b791b26280d', 'Погода'),
								   ('63baeddd-2a07-4f71-aa19-62ecbae26429', 'Курсы валют'),
-- 								   ('', 'Реклама'),
								   ('e6b16a02-3d14-4185-b02b-ef1c3035f159', 'Транспорт'),
								   ('d6e2f387-a6ea-471b-96c3-d46a0e7c796d', 'Время'),
								   ('b71bef49-574e-4354-867c-ca77794172be', 'Новости') ON CONFLICT DO NOTHING;
		`)
	if err != nil {
		panic(err)
	}
	return &Storage{db: db}, nil
}
