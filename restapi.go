package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type product struct {
	ID    string `json:"ID"`
	Name  string `json:"Name"`
	Price string `json:"Price"`
}

var products []product

func createproduct(w http.ResponseWriter, r *http.Request) {

	var newProduct product
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Enter the data correctly")
	}

	json.Unmarshal(reqBody, &newProduct)
	products = append(products, newProduct)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newProduct)
}

func getproducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

func getproduct(w http.ResponseWriter, r *http.Request) {

	productID := mux.Vars(r)["id"]

	for _, product := range products {
		if product.ID == productID {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/products", createproduct).Methods("POST")
	r.HandleFunc("/products", getproducts).Methods("GET")
	r.HandleFunc("/products/{id}", getproduct).Methods("GET")

	fmt.Println("listening at port :8080")
	http.ListenAndServe(":8080", r)
}
