package handlers

import (
	"encoding/json"
	"io"
	"net/http"
)

func Decode(body io.ReadCloser, target interface{}) error {
	return json.NewDecoder(body).Decode(target)
}

func Encode(w http.ResponseWriter, target interface{}) error {
	return json.NewEncoder(w).Encode(target)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}
