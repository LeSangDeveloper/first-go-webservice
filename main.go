package main

import (
	"net/http"
	"pluralsight/gowebservice/product"
)

const apiBasePath = "/api"

func main() {
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}
