package pgsql

import (
	"context"
	"go-core-4/16-db-apps/pkg/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

func (db *DB) Books(ctx context.Context) ([]db.Book, error) {
	return nil, nil
}

func (db *DB) AddBooks(ctx context.Context, books []db.Book) error {
	return nil
}
