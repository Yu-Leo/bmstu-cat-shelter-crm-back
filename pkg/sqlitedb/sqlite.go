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
    photo_url text,
    nickname text NOT NULL, 
    gender text NOT NULL CHECK(gender IN ('male','female')), 
    age integer NOT NULL CHECK(age >= 0),
    chip_number text UNIQUE NOT NULL CHECK(LENGTH(chip_number) == 15) PRIMARY KEY,
    date_of_admission_to_shelter date NOT NULL);`

	_, err := s.DB.ExecContext(ctx, q)
	return err
}
