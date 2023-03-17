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
CREATE TABLE IF NOT EXISTS cats (
    id INTEGER PRIMARY KEY,
    name varchar(80) UNIQUE NOT NULL);`

	_, err := s.DB.ExecContext(ctx, q)
	return err
}
