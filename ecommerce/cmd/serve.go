package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	mux := http.NewServeMux() //router

	// CorsWithPreflight(bruh(logger(mux)))
	// wrappedMux := manager.WrapMux(
	// 	mux,
	// 	middleware.Logger,
	// 	middleware.Bruh,
	// 	middleware.CorsWithPreflight,
	// )

	// var globalMiddlewares []middleware.Middleware
	// globalMiddlewares = append(
	// 	globalMiddlewares,
	// )

	manager.Use(
		middleware.CorsWithPreflight,
		middleware.Bruh,
		middleware.Logger,
	)
	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Server running on :3000")
	err := http.ListenAndServe(":3000", wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
