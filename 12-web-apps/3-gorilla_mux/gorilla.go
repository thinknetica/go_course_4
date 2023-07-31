package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Сторонний маршрутизатор из пакета Gorilla.
	mux := mux.NewRouter()

	// Регистрация обработчика для URL `/` в маршрутизаторе по умолчанию.
	mux.HandleFunc("/{name}", mainHandler).Methods(http.MethodGet)

	// Старт HTTP-сервера на порту 8080 протокола TCP с маршрутизатором запросов "mux".
	http.ListenAndServe(":8080", mux)
}

// HTTP-обработчик.
func mainHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<html><body><h2>Hi, %v</h2></body></html>", vars["name"])
}
