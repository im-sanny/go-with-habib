package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux() //router

	mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("POST /create-product", http.HandlerFunc(createProduct))

	fmt.Println("Server running on :3000")

	globalRouter := globalRouter(mux)

	err := http.ListenAndServe(":3000", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}

func init() {
	prod1 := Product{
		ID:          1,
		Title:       "Fresh Apples",
		Description: "Crisp and sweet red delicious apples",
		Price:       3.99,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRST6XqAF8KkLF0Xqj3vcyfESa4KCCj-Jswhg&s",
	}
	prod2 := Product{
		ID:          2,
		Title:       "Ripe Bananas",
		Description: "Fresh yellow bananas, perfect for smoothies",
		Price:       1.99,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRST6XqAF8KkLF0Xqj3vcyfESa4KCCj-Jswhg&s",
	}
	prod3 := Product{
		ID:          3,
		Title:       "Sweet Oranges",
		Description: "Juicy oranges packed with vitamin C",
		Price:       4.50,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRST6XqAF8KkLF0Xqj3vcyfESa4KCCj-Jswhg&s",
	}

	productList = append(productList, prod1, prod2, prod3)
}

/*
	[] = list
	{} = object
	json = javascript object notation

	If struct field names start with lowercase (e.g. id), they’re unexported and won’t appear in JSON output.

	Keep them uppercase (e.g. ID) and use JSON tags to control key names, like:

	type Product struct {
		ID int `json:"id"`
	}

	The encoding/json package is a separate package, so it can’t access unexported (lowercase) struct fields because they’re private to your package.

	So — lowercase fields = private → invisible to json.
	Uppercase fields = exported → visible to json.

	CORS = cross origin resource sharing
	It’s a security feature in web browsers that controls which websites can request data from your server.


	OPTIONS is an HTTP method.
 	preflightReq = Browser automatically sends a preflight request (using the OPTIONS method) before making certain cross-origin requests. This request checks if the server allows it by verifying the CORS headers in the response.

	SOLID's S is single responsibility principle that means one func should do one particular job, we'll discover more of principle later.

*/
