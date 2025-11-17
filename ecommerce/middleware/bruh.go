package middleware

import (
	"log"
	"net/http"
)

func Bruh(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("I'm bruh middleware")
		next.ServeHTTP(w, r)
	})
}
