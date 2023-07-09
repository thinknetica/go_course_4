package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var api API

func TestAPI_ProductsHandler(t *testing.T) {
	// Создание объекта запроса.
	req := httptest.NewRequest(http.MethodGet, "/api/HANDLER", nil)

	// Создание объекта для записи ответа.
	rr := httptest.NewRecorder()

	// Вызов тестируемого обработчика.
	api.ProductsHandler(rr, req)

	// Проверка ответа.
	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	resp := rr.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var data []product
	json.Unmarshal(body, &data)
	t.Logf("Ответ сервера:\n%v\n", data)
}
