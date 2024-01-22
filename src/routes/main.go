package routes

import (
	"go-api/src/controllers"
	"net/http"
)

func Route(){
	http.HandleFunc("/products", controllers.ProductsController)
	http.HandleFunc("/product/", controllers.ProductController)
}