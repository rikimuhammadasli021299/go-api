package controllers

import (
	"encoding/json"
	"fmt"
	"go-api/src/models"
	"net/http"
)

// get all data and create products
func ProductsController(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET"{
		w.Header().Set("Content-Type", "application/json")
		res := models.SelectAll()
		result , _ := json.Marshal(res.Value)
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else if r.Method == "POST"{
		var input models.Product
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Invalid reqeust body")
			return
		}
		newProduct := models.Product{
			Name: input.Name,
			Price: input.Price,
			Stock: input.Stock,
		}
		res := models.Create(&newProduct)
		var _, _ = json.Marshal(res)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Product Created")
		return
	} else {
		http.Error(w, "", http.StatusBadRequest)
	}
}

// get, update, and delete product by id
func ProductController(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	idParam := r.URL.Path[len("/product/"):]

	if r.Method == "GET" {
		res := models.Select(idParam)	
		result , _ := json.Marshal(res.Value)
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else if r.Method == "PUT" {
		var input models.Product
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Invalid request body")
			return
		}
		
		updateProduct := models.Product{
			Name: input.Name,
			Price: input.Price,
			Stock: input.Stock,
		}
		res := models.Updates(idParam, &updateProduct)
		var _, _ = json.Marshal(res)
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintln(w, "Product Updated")
		return
	} else if r.Method == "DELETE" {
		res := models.Deletes(idParam)
		var _, _ = json.Marshal(res)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Product Deleted")
	} else {
		http.Error(w, "", http.StatusBadRequest)
	}
}