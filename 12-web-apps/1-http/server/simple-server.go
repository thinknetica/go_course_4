package main

import (
	"net/http"
)

func main() {
	// регистрация обработчика для URL `/` в маршрутизаторе по умолчанию
	http.HandleFunc("/", mainHandler)

	// старт HTTP-сервера на порту 8080 протокола TCP с маршрутизатором запросов по умолчанию
	http.ListenAndServe(":8080", nil)
}

// HTTP-обработчик
func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<html><body><h2>Go Simple Web App</h2></body></html>`))
}
