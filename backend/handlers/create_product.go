package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Product Data", 400)
		return
	}
	newProduct.ID = len(database.ProductList) + 1 // Assign a new ID
	database.ProductList = append(database.ProductList, newProduct)

	util.SendData(w, newProduct, 201)
}
