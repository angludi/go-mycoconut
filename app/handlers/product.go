package handlers

import (
	"net/http"
	"my-coconut.com/web/app/models"
	"encoding/json"
	"my-coconut.com/web/app/transformer"
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

	res := new(transformer.ProdukCollectionTransformer)
	res.Transform(produk)
	RespondJson(w, http.StatusOK, res)
}