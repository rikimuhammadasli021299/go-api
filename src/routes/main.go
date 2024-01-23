package routes

import (
	"go-api/src/controllers"
	"go-api/src/middleware"
	"net/http"
)

func Route(){
	http.HandleFunc("/register", controllers.RegisterUser)
	http.HandleFunc("/login", controllers.Login)
	http.Handle("/products", middleware.JwtMiddleware(http.HandlerFunc(controllers.ProductsController)))
	http.Handle("/product/", middleware.JwtMiddleware(http.HandlerFunc(controllers.ProductController)))
}