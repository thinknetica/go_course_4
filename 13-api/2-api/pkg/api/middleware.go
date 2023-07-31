package api

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

// проверка валидности сеанса
func (api *API) sessionsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := api.store.Get(r, "session-cookie")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "доступ запрещён", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// проверка валидности JWT
func (api *API) jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			next.ServeHTTP(w, r)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenString := splitted[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret-password"), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Printf("Данные токена JWT: %+v\n", claims)
		}

		next.ServeHTTP(w, r)
	})
}

// присвоение номера запросу
func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "request_id", rand.Intn(1_000_000))
		newR := r.WithContext(ctx)
		b, _ := httputil.DumpRequest(newR, true)
		fmt.Printf("%+v", string(b))

		next.ServeHTTP(w, newR)
	})
}

// логирование запросов к API
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value("request_id").(int)
		log.Println(r.Method, r.RemoteAddr, r.RequestURI, id)
		next.ServeHTTP(w, r)
	})
}
