package handlers

import (
	"encoding/json"
	"io"
)

func Decode(body io.ReadCloser, target interface{}) error {
	return json.NewDecoder(body).Decode(target)
}

func Encode(w io.Writer, target interface{}) error {
	return json.NewEncoder(w).Encode(target)
}
