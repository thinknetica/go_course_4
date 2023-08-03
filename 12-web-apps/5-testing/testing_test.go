package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

var testMux *mux.Router

func TestMain(m *testing.M) {
	testMux = mux.NewRouter()
	endpoints(testMux)
	m.Run()
}

func Test_mainHandler(t *testing.T) {
	data := []int{}
	payload, _ := json.Marshal(data)

	// Создаём HTTP=запрос.
	req := httptest.NewRequest(http.MethodPost, "/Name", bytes.NewBuffer(payload))
	req.Header.Add("Сontent-type", "plain/text")

	// Объект для записи ответа HTTP-сервера.
	rr := httptest.NewRecorder()

	// Вызов маршрутизатора и обслуживание запроса.
	testMux.ServeHTTP(rr, req)

	// Анализ ответа сервера (неверный метод HTTP).
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	t.Log("Response: ", rr.Body)

	//=========================================================

	req = httptest.NewRequest(http.MethodGet, "/Name", nil)
	req.Header.Add("Сontent-type", "plain/text")

	rr = httptest.NewRecorder()

	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	body := rr.Body.String()
	if !strings.Contains(body, "Name") {
		t.Fatal(body)
	}
}
