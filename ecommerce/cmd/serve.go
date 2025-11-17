package cmd

import (
	"ecommerce/globalRouter"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	manager.Use(middleware.Logger, middleware.Bruh)

	mux := http.NewServeMux() //router

	initRoutes(mux, manager)

	globalRouter := globalRouter.GlobalRouter(mux)

	fmt.Println("Server running on :3000")
	err := http.ListenAndServe(":3000", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
