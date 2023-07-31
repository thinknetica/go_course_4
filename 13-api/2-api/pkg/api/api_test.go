package api

import (
	"os"
	"testing"

	"github.com/gorilla/mux"
)

var api *API

func TestMain(m *testing.M) {
	router := mux.NewRouter()
	api = New(router)
	api.Endpoints()
	os.Exit(m.Run())
}
