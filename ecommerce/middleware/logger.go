package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Println("I'm middleware: I'll be print first")

		next.ServeHTTP(w, r)

		log.Println("I'm middleware, I'll be print later")

		log.Println(r.Method, r.URL.Path, time.Since(start))
	})

}
