package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPI_jwtMiddleware(t *testing.T) {
	data := authInfo{
		Usr: "Usr",
		Pwd: "Pwd",
	}
	payload, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/authJWT", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	b, _ := ioutil.ReadAll(rr.Body)
	jwt := string(b)

	req = httptest.NewRequest(http.MethodGet, "/api/v1/books", nil)
	req.Header.Add("Authorization", "Bearer "+jwt)
	rr = httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
}
