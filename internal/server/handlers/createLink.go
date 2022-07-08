package handlers

import (
	"net/http"

	"github.com/gsmerlin/minify/internal/db"
	"github.com/gsmerlin/minify/internal/logger"
)

type CreateSchema struct {
	Destination string `json:"destination"`
	Email       string `json:"email"`
}

func CreateLink(w http.ResponseWriter, r *http.Request) {

	var json CreateSchema

	if err := getJson(r.Body, &json); err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.Info("Creating link for " + json.Email + " to " + json.Destination)

	id := db.NewLink("", json.Email, json.Destination)

	w.Write([]byte(id))
}
