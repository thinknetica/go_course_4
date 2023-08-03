package main

import (
	"net/http"

	"go-core-4/13-api/2-api/pkg/api"
)

type server struct {
	api *api.API
}

func main() {
	srv := new(server)

	srv.api = api.New()

	http.ListenAndServe(":8082", srv.api.Router())
}
