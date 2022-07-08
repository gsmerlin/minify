package handlers

import (
	"encoding/json"
	"io"
)

func getJson(body io.ReadCloser, target interface{}) error {
	return json.NewDecoder(body).Decode(target)
}
