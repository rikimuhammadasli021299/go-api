package main

import (
	"fmt"
	"go-api/src/config"
	"go-api/src/helper"
	"go-api/src/routes"
	"net/http"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.InitDB()
	helper.Migration()
	defer config.DB.Close()
	routes.Route()
	fmt.Println("server is running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}