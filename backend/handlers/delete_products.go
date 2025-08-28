package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func DeleteProducts(w http.ResponseWriter, r *http.Request) {
	pId := r.PathValue("productId")
	delId, err := strconv.Atoi(pId)
	if err != nil {
		util.SendData(w, "Invalid Product ID", 400)
		return
	}
	for i := range database.ProductList {
		if database.ProductList[i].ID == delId {
			database.ProductList = append(database.ProductList[:i], database.ProductList[i+1:]...)
			util.SendData(w, "Product Deleted Seccessfully!", 200)
			return
		}
	}
	util.SendData(w, "Product Not Found", 404)
}
