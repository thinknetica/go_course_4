package testmain

import (
	"context"
	"log"
	"reflect"
	"testing"

	"github.com/jackc/pgx/v4"
)

var testService *Service

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(context.Background(), "postgres://user:pwd@server/database")
	if err != nil {
		log.Println(err)
	}
	testService = &Service{
		db: conn,
	}
	m.Run()
	// Здесь может идти освобождение ресурсов.
}

func TestService_Products(t *testing.T) {
	got := testService.Products()
	want := []Product{
		{
			Name:  "Компьютер",
			Price: 20_000,
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Service.Products() = %v, want %v", got, want)
	}
}
