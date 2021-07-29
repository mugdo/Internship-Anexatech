package main

import (
	"fmt"
	"net/http"
	"quickstart/vapi"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Project Start...")

	r := mux.NewRouter()

	r.HandleFunc("/api/product/{id}", vapi.Getproduct).Methods("GET")
	r.HandleFunc("/api/products", vapi.Getproducts).Methods("GET")
	r.HandleFunc("/api/product", vapi.Addproduct).Methods("POST")
	r.HandleFunc("/api/product/{id}", vapi.Deleteproduct).Methods("DELETE")
	r.HandleFunc("/api/product/{id}", vapi.Updateproduct).Methods("PUT")
	r.HandleFunc("/api/product/category/{value}", vapi.Productcategory).Methods("GET")

	http.ListenAndServe(":8000", r)
}
