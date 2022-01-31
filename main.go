package main

import (
	"net/http"
	"pluralsight/gowebservice/database"
	"pluralsight/gowebservice/product"

	_ "github.com/go-sql-driver/mysql"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}
