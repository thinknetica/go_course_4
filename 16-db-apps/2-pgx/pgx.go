package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Книжка.
type book struct {
	ID          int
	Title       string
	Year        int
	Public      bool
	PublisherID int
	Publisher   string
}

func main() {
	ctx := context.Background()
	// Подключение к БД. Функция возвращает объект БД.
	pwd := os.Getenv("pg_password")
	db, err := pgxpool.New(ctx, "postgres://postgres:"+pwd+"@ubuntu-server.northeurope.cloudapp.azure.com/books")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	// Не забываем очищать ресурсы.
	defer db.Close()

	// Получение списка книг.
	data, err := books(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)

	// добавление книг одной транзакцией
	data = []book{
		{
			Title: "The Chronicles of Amber",
			Year:  1970,
		},
		{
			Title: "Dune",
			Year:  1965,
		},
	}
	err = addBooks(ctx, db, data)
	if err != nil {
		log.Fatal(err)
	}

	// Получение списка книг.
	data, err = books(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
}

// books возвращает список книг из БД.
func books(ctx context.Context, db *pgxpool.Pool) ([]book, error) {
	// Выполнение запроса выборки данных.
	// Query производит следующие действия:
	// - подготовка запроса;
	// - выполнение запроса;
	// - закрытие соединения.
	rows, err := db.Query(ctx, `
		SELECT id, title, year
		FRoM books
		WHERE id >= $1`,
		0,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []book
	for rows.Next() {
		var b book
		err := rows.Scan(
			&b.ID,
			&b.Title,
			&b.Year,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return books, nil
}

// addBooks добавляет в БД массив книг одной транзакцией.
func addBooks(ctx context.Context, db *pgxpool.Pool, books []book) error {
	// начало транзакции
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	// отмена транзакции в случае ошибки
	defer tx.Rollback(ctx)

	// пакетный запрос
	var batch = &pgx.Batch{}
	// добавление заданий в пакет
	for _, book := range books {
		batch.Queue(`INSERT INTO books(title, year) VALUES ($1, $2)`, book.Title, book.Year)
	}
	// отправка пакета в БД (может выполняться для транзакции или соединения)
	res := tx.SendBatch(ctx, batch)
	// обязательная операция закрытия соединения
	err = res.Close()
	if err != nil {
		return err
	}
	// подтверждение транзакции
	return tx.Commit(ctx)
}
