package cmd

import (
	"ecommerce/globalRouter"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux() //router

	mux.Handle("GET /name", middleware.Logger(http.HandlerFunc(handlers.Test)))

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductsByID))

	globalRouter := globalRouter.GlobalRouter(mux)

	fmt.Println("Server running on :3000")
	err := http.ListenAndServe(":3000", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
