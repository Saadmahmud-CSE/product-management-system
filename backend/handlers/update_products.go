package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdateProducts(w http.ResponseWriter, r *http.Request) {
	pId := r.PathValue("productId")
	upId, err := strconv.Atoi(pId)
	if err != nil {
		util.SendData(w, "Invalid Product ID", 400)
		return
	}

	// Decode new item data from request body
	var updatedData database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedData)
	if err != nil {
		util.SendData(w, "Invalid Request Body", 400)
		return
	}

	// Find and update the item
	for i := range database.ProductList {
		if database.ProductList[i].ID == upId {
			database.ProductList[i].Title = updatedData.Title
			database.ProductList[i].Price = updatedData.Price
			database.ProductList[i].Description = updatedData.Description
			database.ProductList[i].Category = updatedData.Category
			database.ProductList[i].ImgUrl = updatedData.ImgUrl
			database.ProductList[i].Rating.Rate = updatedData.Rating.Rate
			database.ProductList[i].Rating.Count = updatedData.Rating.Count

			util.SendData(w, database.ProductList[i], 200)
			return
		}
	}

	util.SendData(w, "Product Not Found", 404)
}
