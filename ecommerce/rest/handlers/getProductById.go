package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductsByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productId")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	product := database.Get(pId)
	if product == nil {
		util.SendError(w, 404, "Product not found")
		return
	}

	util.SendData(w, product, 200)

}
