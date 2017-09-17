package main

import (
	"github.com/gorilla/mux"
	"my-coconut.com/web/app/handlers"
	"my-coconut.com/web/config"
	"net/http"
	"fmt"
	"log"
)

var appConfig = config.Config.App

func main() {
	r := mux.NewRouter()

	product := r.PathPrefix("/produk").Subrouter()
	product.HandleFunc("", handlers.GetAllProducts).Methods("GET")

	kategori := r.PathPrefix("/kategori").Subrouter()
	kategori.HandleFunc("", handlers.GetAllKategori).Methods("GET")

	err := http.ListenAndServe(fmt.Sprintf(":%s", appConfig.HttpPort), &MyServer{r})
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

type MyServer struct {
	r *mux.Router
}

func (s *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	s.r.ServeHTTP(w, r)
}