package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// сторонний маршрутизатор из пакета Gorilla
var r *mux.Router

func main() {
	r = mux.NewRouter()
	endpoints(r)
	http.ListenAndServe(":8080", r)
}

func endpoints(r *mux.Router) {
	// Регистрация обработчика для URL `/` в маршрутизаторе по умолчанию.
	r.HandleFunc("/{name}", mainHandler).Methods(http.MethodGet)
}

// HTTP-обработчик.
func mainHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<html><body><h2>Hi, %v</h2></body></html>", vars["name"])
}
