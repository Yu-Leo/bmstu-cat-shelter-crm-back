package sqliteStorage

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(path string) *Storage {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if err := db.Ping(); err != nil {
		fmt.Println(err)
		return nil
	}
	return &Storage{DB: db}
}

func (s *Storage) Init(ctx context.Context) error {
	q := `CREATE TABLE IF NOT EXISTS cats (
    id INTEGER PRIMARY KEY,
    name varchar(80) UNIQUE NOT NULL);`

	_, err := s.DB.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}

	return nil
}
