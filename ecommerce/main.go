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
	w.Header().Set("Access-Control-Allow-Origin", "*") // anyone can now access this api
	w.Header().Set("Content-Type", "application/json") //sets the response type to JSON, so the client knows how to read and parse it properly.

	if r.Method != http.MethodGet { //r.Method = post, put, patch, delete
		http.Error(w, "please give me GET request", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

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
	w.WriteHeader(201)
	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
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

*/
