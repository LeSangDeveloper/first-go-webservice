package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"pluralsight/gowebservice/cors"
	"strconv"
	"strings"
)

const productsBasePath = "products"

func SetupRoutes(apiBasePath string) {
	handleProducts := http.HandlerFunc(productsHandler)
	handleProduct := http.HandlerFunc(productHandler)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, productsBasePath), cors.Middleware(handleProducts))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, productsBasePath), cors.Middleware(handleProduct))
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegment := strings.Split(r.URL.Path, "products/")
	productID, err := strconv.Atoi(urlPathSegment[len(urlPathSegment)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		productObject, err := getProduct(productID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if productObject == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		productJSON, err := json.Marshal(productObject)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productJSON)
	case http.MethodPut:
		var updatedProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &updatedProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updatedProduct.ProductID != productID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = updateProduct(updatedProduct)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodDelete:
		removeProduct(productID)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productList, err := getProductList()
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		productJson, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productJson)
	case http.MethodPost:
		var newProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(bodyBytes, &newProduct)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if newProduct.ProductID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = insertProduct(newProduct)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	case http.MethodOptions:
		return
	}
}
