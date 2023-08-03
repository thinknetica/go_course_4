package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func (api *API) authJWT(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var auth authInfo
	err = json.Unmarshal(body, &auth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if auth.Usr == "Usr" && auth.Pwd == "Pwd" {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"usr": auth.Usr,
			"nbf": time.Now().Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString([]byte("secret-password"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(tokenString))
	}
}
