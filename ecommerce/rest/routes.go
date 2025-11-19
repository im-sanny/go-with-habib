package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /fumis", manager.With(
		http.HandlerFunc(handlers.Test),
	))

	mux.Handle("GET /name", manager.With(
		http.HandlerFunc(handlers.Test),
	))

	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProducts),
	))

	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(handlers.CreateProduct),
	))

	mux.Handle("GET /products/{productId}", manager.With(
		http.HandlerFunc(handlers.GetProductsByID),
	))

}
