package main

import (
	"net/http"

	"go-core-4/13-api/2-api/pkg/api"

	"github.com/gorilla/mux"
)

type server struct {
	api    *api.API
	router *mux.Router
}

func main() {
	srv := new(server)
	srv.router = mux.NewRouter()
	srv.api = api.New(srv.router)
	srv.api.Endpoints()
	http.ListenAndServe(":8082", srv.router)
}
