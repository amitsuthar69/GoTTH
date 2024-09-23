package middleware

import (
	"log"
	"net/http"
)

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Authorization")
		h.ServeHTTP(w, r)
	})
}
