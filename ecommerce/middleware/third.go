package middleware

import (
	"log"
	"net/http"
)

func Third(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("I'm third middleware")
		next.ServeHTTP(w, r)
	})
}
