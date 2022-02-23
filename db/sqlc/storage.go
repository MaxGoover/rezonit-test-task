package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Storage provides all functions to execute db queries and transaction
type Storage struct {
	db *sql.DB
	*Queries
}

// NewStorage creates a new store
func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		db:      db,
		Queries: New(db),
	}
}

// ExecTx executes a function within a database transaction
func (store *Storage) execTx(ctx context.Context, fn func(*Queries) error) error {
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
