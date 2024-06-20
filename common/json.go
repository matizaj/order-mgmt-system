package common

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ErrorJson(w http.ResponseWriter, status int, msg string) {
	WriteJson(w, status, map[string]string{"error":msg})
}

func ReadJson(r *http.Request, data any) error {
	return json.NewDecoder(r.Body).Decode(data)
}