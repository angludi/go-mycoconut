package handlers

import (
	"net/http"
	"my-coconut.com/web/app/models"
	"encoding/json"
)

func GetAllProducts (w http.ResponseWriter, r *http.Request) {
	produk, err := models.GetProducts()

	if err != nil {
		resp, err := json.Marshal(err)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(resp)
		return
	}

	RespondJson(w, http.StatusOK, produk)
}