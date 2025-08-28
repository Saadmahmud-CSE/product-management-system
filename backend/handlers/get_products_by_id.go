package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productId")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		panic(err)
	}
	for _, product := range database.ProductList {
		if product.ID == pId {
			util.SendData(w, product, 200)
			return
		}
	}
	util.SendData(w, "Product not found", 404)
}
