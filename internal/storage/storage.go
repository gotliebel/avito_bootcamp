package storage

import (
	"avito_bootcamp/internal/constants"
	"database/sql"
	_ "github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Storage struct {
	Db *sql.DB
}

func NewStorage() (*Storage, error) {
	db, err := sql.Open("pgx", constants.DataBaseConnection)
	if err != nil {
		return nil, err
	}
	return &Storage{db}, nil
}

func (s *Storage) Close() {
	s.Db.Close()
}
