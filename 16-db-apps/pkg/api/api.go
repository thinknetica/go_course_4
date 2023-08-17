package api

import (
	"encoding/json"
	"go-core-4/16-db-apps/pkg/db"
	"net/http"
)

type API struct {
	db db.Interface
}

func New(db db.Interface) *API {
	return &API{db: db}
}

func (api *API) handler(w http.ResponseWriter, r *http.Request) {
	books, _ := api.db.Books(r.Context())

	for _, book := range books {
		_ = book
	}

	json.NewEncoder(w).Encode(books)
}
