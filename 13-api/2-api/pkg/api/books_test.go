package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPI_newBook(t *testing.T) {
	want := len(books) + 1 // что здесь плохо?
	data := book{
		Name:   "1984",
		Author: "George Orwell",
	}

	payload, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/newBook", bytes.NewBuffer(payload))

	rr := httptest.NewRecorder()

	api.router.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	t.Log("Response: ", rr.Body)

	got := len(books)
	// что плохо в таком сравнении? как сделать лучше?
	if got != want {
		t.Fatal("книга не добавлена")
	}
}

func TestAPI_books(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/books", nil)
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	t.Log("Response: ", rr.Body)
}
