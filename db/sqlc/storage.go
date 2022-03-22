package db

import (
	"database/sql"
)

type Storage interface {
	Querier
}

type SQLStorage struct {
	db *sql.DB
	*Queries
}

func NewStorage(db *sql.DB) Storage {
	return &SQLStorage{
		db:      db,
		Queries: New(db),
	}
}
