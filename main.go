package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type product struct{
	Id    int	 `json:"id"`
	Name  string `json:"name"`
	Price int	 `json:"price"`
	Stock int	 `json:"stock"`
}

var products = []product{
	product{1,"T-Shirt",100000,12},
	product{2,"Pants",80000,20},
	product{3,"Hoodie",150000,20},
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/products", productsController)
	http.HandleFunc("/product/", productController)
	fmt.Println("server is running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func sayHello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Hello World!")
}

// get all data products
func productsController(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	result , _ := json.Marshal(products)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

// get data products by id
func productController(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	idParam := r.URL.Path[len("/product/"):]
	id, _ := strconv.Atoi(idParam)
	
	var foundIndex = -1
	for i, p := range products{
		if p.Id == id{
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	}

	result , _ := json.Marshal(products[foundIndex])
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}