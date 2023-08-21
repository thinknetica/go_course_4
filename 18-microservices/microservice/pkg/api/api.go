package api

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
)

type API struct {
	Router *mux.Router

	sync.Mutex
	db []Review
}

func New() *API {
	api := API{
		Router: mux.NewRouter(),
	}

	api.Router.HandleFunc("/", handler)

	return &api
}

type Review struct {
	MocieID int
	Text    string
}

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	w.Write([]byte(fmt.Sprintf("Microapp на машине %v", hostname)))
}
