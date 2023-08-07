package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/rs/zerolog"
)

// API предоставляет интерфейс программного взаимодействия.
type API struct {
	router *mux.Router
	store  *sessions.CookieStore

	log *zerolog.Logger
}

// New создаёт объект API.
func New() *API {
	api := API{
		router: mux.NewRouter(),
		store:  sessions.NewCookieStore([]byte("secret_password")),
	}

	api.endpoints()

	return &api
}

func (api *API) Router() *mux.Router {
	return api.router
}

// endpoints регистрирует конечные точки API.
func (api *API) endpoints() {
	api.router.Use(requestIDMiddleware)
	api.router.Use(logMiddleware)
	api.router.Use(api.jwtMiddleware)
	//api.router.Use(api.sessionsMiddleware)

	api.router.HandleFunc("/api/v1/authSession", api.authSession).Methods(http.MethodPost, http.MethodOptions)
	api.router.HandleFunc("/api/v1/authJWT", api.authJWT).Methods(http.MethodPost, http.MethodOptions)

	api.router.HandleFunc("/api/v1/books", api.books).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/api/v1/newBook", api.newBook).Methods(http.MethodPost, http.MethodOptions)

	rr, _ := http.NewRequestWithContext(context.Background(), "GET", "https://yandex.ru", nil)

	http.DefaultClient.Do(rr)
}
