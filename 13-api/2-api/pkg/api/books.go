package api

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

func (api *API) books(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) newBook(w http.ResponseWriter, r *http.Request) {
	// контекст с таймаутом
	timeout, cancel := context.WithTimeout(r.Context(), time.Second*10)
	// освобождение ресурсов, если функция завершится раньше
	defer cancel()
	// производный контекст с ключом (№ запроса)
	ctx := context.WithValue(timeout, "requestID", rand.Intn(1_000_000_000))
	// длительная операция, принимающая контекст
	//data := getLotsOfDataFromDatabase(ctx)
	_ = ctx

	var b book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	books = append(books, b)
}

type book struct {
	Name   string
	Author string
}

var books = []book{
	{
		Name:   "The Lord Of The Rings",
		Author: "J.R.R. Tolkien",
	},
}
