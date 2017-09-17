package main

import (
	"github.com/gorilla/mux"
	"my-coconut.com/web/app/handlers"
)

func main() {
	r := mux.NewRouter()

	product := r.PathPrefix("/products").Subrouter()
	product.HandleFunc("", )
}


