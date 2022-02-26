package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Storage interface {
	Querier
}

// SQLStorage provides all functions to execute SQL queries and transaction
type SQLStorage struct {
	db *sql.DB
	*Queries
}

// NewStorage creates a new store
func NewStorage(db *sql.DB) Storage {
	return &SQLStorage{
		db:      db,
		Queries: New(db),
	}
}

// ExecTx executes a function within a database transaction
func (store *SQLStorage) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
