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
	if r.Method == "GET"{
		w.Header().Set("Content-Type", "application/json")
		result , _ := json.Marshal(products)
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else if  r.Method == "POST"{
		var product product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Invalid reqeust body")
			return
		}
		products = append(products, product)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Product Created")
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

// get, create, update, delete data products by id
func productController(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	idParam := r.URL.Path[len("/product/"):]
	id, _ := strconv.Atoi(idParam)

	// Check id
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

	if r.Method == "GET" {	
		result , _ := json.Marshal(products[foundIndex])
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else if r.Method == "PUT" {
		var updateProduct product

		err := json.NewDecoder(r.Body).Decode(&updateProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Invalid request body")
			return
		}

		products[foundIndex] = updateProduct
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintln(w, "Product Updated")
		return
	} else if r.Method == "DELETE" {
		_ = append(products[:foundIndex], products[foundIndex+1:]...)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Product Deleted")
	}
	http.Error(w, "", http.StatusBadRequest)
}