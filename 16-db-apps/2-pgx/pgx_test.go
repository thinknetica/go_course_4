package main

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	testDB *pgxpool.Pool
	ctx    = context.Background()
)

func TestMain(m *testing.M) {
	var err error
	// БД для тетов.
	testDB, err = pgxpool.New(context.Background(), "postgres://postgres:GoPassword@ubuntu-server.northeurope.cloudapp.azure.com/books")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer testDB.Close()
	os.Exit(m.Run())
}

func Test_books(t *testing.T) {
	data, err := books(ctx, testDB)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", data)
}
