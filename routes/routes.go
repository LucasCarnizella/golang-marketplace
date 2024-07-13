package routes

import (
	"marketplace/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/product/new", controllers.NewProductPage)
	http.HandleFunc("/product/edit", controllers.EditProductPage)
	http.HandleFunc("/product/create", controllers.CreateProduct)
	http.HandleFunc("/product/update", controllers.UpdateProduct)
	http.HandleFunc("/product/delete", controllers.DeleteProduct)
}
