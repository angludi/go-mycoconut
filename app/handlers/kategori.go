package handlers

import (
	"net/http"
	"my-coconut.com/web/app/models"
	"encoding/json"
)

func GetAllKategori (w http.ResponseWriter, r *http.Request) {
	produkKategori, err := models.GetKategori()

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

	RespondJson(w, http.StatusOK, produkKategori)
}