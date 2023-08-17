package api

import (
	"encoding/json"
	"go-core-4/17-system-design/app/internal/db"
	"go-core-4/17-system-design/app/internal/models"
	"net/http"

	"github.com/gorilla/mux"
)

// API предоставляет интерфейс программного взаимодействия.
type API struct {
	router *mux.Router
	db     *db.DB
}

func New(db *db.DB) *API {
	api := API{
		router: mux.NewRouter(),
		db:     db,
	}
	api.Endpoints()
	return &api
}

func (api *API) Run(addr string) error {
	return http.ListenAndServe(addr, api.router)
}

// Endpoints регистрирует конечные точки API.
func (api *API) Endpoints() {
	api.router.HandleFunc("/api/v1/books", api.books).Methods(http.MethodGet)
	api.router.HandleFunc("/api/v1/books", api.newBook).Methods(http.MethodPost)
	api.router.HandleFunc("/api/v1/books/{id}", api.deleteBook).Methods(http.MethodDelete)
}

func (api *API) books(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(models.Books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) newBook(w http.ResponseWriter, r *http.Request) {
	var b models.Book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	models.Books = append(models.Books, b)
}

func (api *API) deleteBook(w http.ResponseWriter, r *http.Request) {
}
