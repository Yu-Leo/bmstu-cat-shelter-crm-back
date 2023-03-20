package sqlitedb

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Storage{DB: db}, nil
}

func (s *Storage) Init(ctx context.Context) error {
	q := `
DROP TABLE IF EXISTS cats;
CREATE TABLE cats (
    id integer PRIMARY KEY,
    photo_url varchar(80),
    nickname varchar(80) UNIQUE NOT NULL, 
    gender boolean NOT NULL, 
    age integer NOT NULL,
    chip_number varchar(15) UNIQUE NOT NULL,
    date_of_admission_to_shelter date NOT NULL);`
	/*
		TODO:
			photo_url: ?
			nickname: 80 ?
			gender: Enum (?), check
			age: not negative
			chip_number: unique ?, constraint: only digits
	*/

	_, err := s.DB.ExecContext(ctx, q)
	return err
}
