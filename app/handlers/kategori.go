package handlers

import (
	"net/http"
	"my-coconut.com/web/app/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func GetAllKategori (w http.ResponseWriter, r *http.Request) {
	produkKategori, err := models.GetAllKategori()

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

func GetKategori(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10,32)

	kategori, err := models.GetKategoriByProductID(uint32(id))

	if err != nil {
		RespondError(w, "Cannot get admin", http.StatusUnauthorized)
	}

	RespondJson(w, http.StatusOK, kategori)
}