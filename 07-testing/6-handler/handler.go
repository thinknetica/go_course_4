package handler

import (
	"encoding/json"
	"net/http"
)

// API - API.
type API struct {
	dbConn string
}

type product struct {
	Name  string
	Price int
}

func (api API) ProductsHandler(w http.ResponseWriter, r *http.Request) {
	data := []product{
		{
			Name:  "Книга",
			Price: 1000,
		},
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
