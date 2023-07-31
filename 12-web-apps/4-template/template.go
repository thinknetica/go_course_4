package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/{first}/{last}", mainHandler).Methods(http.MethodGet)

	http.ListenAndServe(":8080", mux)
}

// HTTP-обработчик, возвращающий HTML-страницу.
func mainHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := template.New("main")
	t, err := t.Parse("<html><body><h2>Hi, {{.}}</h2></body></html>")
	if err != nil {
		http.Error(w, "ошибка при обработке шаблона", http.StatusInternalServerError)
		return
	}
	t.Execute(w, vars["first"]+vars["last"])
}
