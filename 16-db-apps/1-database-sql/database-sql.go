package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Использование шаблона "Репозиторий".
type BooksRepo interface {
	GetBooks(context.Context) ([]book, error)
	AddBook(context.Context, book) (int, error)
}

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
	// Объект БД - пул подключений к СУБД.
	// БД - долгоживущий объект. Следует создавать только один объект для каждой БД.
	// Далее этот объект следует передавать как зависимость.
	var db *sql.DB
	var err error

	// Подключение к БД.
	// В зависимости от драйвера, sql.Open может не выполнять фактического подключения,
	// а только проверить параметры соединения с БД.
	pwd := os.Getenv("mysql_password")
	db, err = sql.Open("mysql", "root:"+pwd+"@tcp(ubuntu-server.northeurope.cloudapp.azure.com:3306)/books")
	if err != nil {
		log.Fatal(err)
	}
	// Не забываем очищать ресурсы.
	defer db.Close()

	// Проверка соединения с БД. На случай, если sql.Open этого не делает.
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Получение списка книг.
	data, err := books(db)
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
	err = addBooks(db, data)
	if err != nil {
		log.Fatal(err)
	}

	// Получение списка книг.
	data, err = books(db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)
}

// books возвращает список книг из БД.
func books(db *sql.DB) ([]book, error) {
	// Выполнение запроса выборки данных.
	// Query производит следующие действия:
	// - подготовка запроса;
	// - выполнение запроса;
	// - закрытие соединения.
	rows, err := db.Query(`
		SELECT id, title, year
		FROM books
		WHERE id >= ?`,
		0,
	)
	if err != nil {
		return nil, err
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
func addBooks(db *sql.DB, books []book) error {
	// начало транзакции
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// отмена транзакции в случае ошибки
	defer tx.Rollback()

	// подготовка запроса для последующего многократного выполнения
	stmt, err := tx.Prepare(`INSERT INTO books(title, year) VALUES (?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, book := range books {
		res, err := stmt.Exec(book.Title, book.Year)
		if err != nil {
			return err
		}
		id, _ := res.LastInsertId()
		fmt.Println("Создана запись с ID:", id)
	}

	// подтверждение транзакции
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
