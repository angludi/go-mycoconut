package handlers

import (
"encoding/json"
"net/http"
)

type ErrorResponse struct {
	Message string `json:"errors"`
}

func RespondJson(w http.ResponseWriter, statusCode int, i interface{}) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(resp)
}

func RespondError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	erresp := ErrorResponse{message}
	resp, err := json.Marshal(erresp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Write(resp)
}

func RespondFailValidation(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	erresp := ErrorResponse{message}
	resp, err := json.Marshal(erresp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(resp)
}

