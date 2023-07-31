package main

import (
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	const addr = ":8080"

	// Параметры веб-сервера.
	srv := &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		Handler:      nil,
		Addr:         addr,
	}

	// Старт сетевой службы веб-сервера.
	listener, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Fatal(err)
	}

	// Регистрация обработчика для URL `/` в маршрутизаторе по умолчанию.
	http.HandleFunc("/", mainHandler)

	// Старт самого веб-сервера.
	log.Fatal(srv.Serve(listener))
	//log.Fatal(srv.ServeTLS(listener, "cert.crt", "cert.key"))
}

// HTTP-обработчик
func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<html><body><h2>Go Simple Web App</h2></body></html>`))
}
