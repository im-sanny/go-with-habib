package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm a software engineer")
}

type Product struct {
	ID          int
	Title       string
	Description string
	Price       float64
	ImgUrl      string
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightReq(w, r)

	if r.Method != http.MethodGet { //r.Method = post, put, patch, delete
		http.Error(w, "please give me GET request", 400)
		return
	}

	sendData(w, productList, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightReq(w, r)

	if r.Method != "POST" { //r.Method = post, put, patch, delete
		http.Error(w, "please give me POST request", 400)
		return
	}

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

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Name") // for this to work i need to set a custom header from frontend code
	w.Header().Set("Content-Type", "application/json")
}

func handlePreflightReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(200)
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux() //router

	mux.HandleFunc("/hello", helloHandler) //route
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/create-product", createProduct)

	fmt.Println("Server running on :3000")

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}

func init() {
	prod1 := Product{
		ID:          1,
		Title:       "Fruit 1",
		Description: "Fruits are delicious",
		Price:       5,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRST6XqAF8KkLF0Xqj3vcyfESa4KCCj-Jswhg&s",
	}
	prod2 := Product{
		ID:          2,
		Title:       "Fruit 1",
		Description: "Fruits are delicious",
		Price:       5,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRST6XqAF8KkLF0Xqj3vcyfESa4KCCj-Jswhg&s",
	}
	prod3 := Product{
		ID:          3,
		Title:       "Fruit 1",
		Description: "Fruits are delicious",
		Price:       5,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRST6XqAF8KkLF0Xqj3vcyfESa4KCCj-Jswhg&s",
	}
	prod4 := Product{
		ID:          4,
		Title:       "Fruit 1",
		Description: "Fruits are delicious",
		Price:       5,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRST6XqAF8KkLF0Xqj3vcyfESa4KCCj-Jswhg&s",
	}
	prod5 := Product{
		ID:          5,
		Title:       "Fruit 1",
		Description: "Fruits are delicious",
		Price:       5,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRST6XqAF8KkLF0Xqj3vcyfESa4KCCj-Jswhg&s",
	}

	productList = append(productList, prod1, prod2, prod3, prod4, prod5)
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
