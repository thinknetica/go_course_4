package db

import (
	"context"
)

type Interface interface {
	Books(ctx context.Context) ([]Book, error)
	AddBooks(ctx context.Context, books []Book) error
}

type Book struct {
	ID          int
	Title       string
	Year        int
	Public      bool
	PublisherID int
	Publisher   string
}
