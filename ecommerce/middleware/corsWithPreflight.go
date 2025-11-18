package middleware

import (
	"log"
	"net/http"
)

func CorsWithPreflight(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("I'm cors middleware")
		// handle cors
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Method", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Name") // for this to work i need to set a custom header from frontend code
		w.Header().Set("Content-Type", "application/json")

		// handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		}
		next.ServeHTTP(w, r)

	})
}
