package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int
	Title       string
	Description string
	Price       float64
	ImgUrl      string
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	sendData(w, productList, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "please give me valid json", 400)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)
	sendData(w, newProduct, 201)
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

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

func globalRouter(mux *http.ServeMux) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Method", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Name") // for this to work i need to set a custom header from frontend code
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		}
		mux.ServeHTTP(w, r)

	}
	return http.HandlerFunc(handleAllReq)
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
