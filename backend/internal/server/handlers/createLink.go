package handlers

import (
	"net/http"

	"github.com/gsmerlin/minify/backend/internal/db"
	"github.com/gsmerlin/minify/backend/internal/logger"
)

type CreateLinkInput struct {
	Destination string `json:"destination"`
	Email       string `json:"email"`
}

type CreateLinkOutput struct {
	ID string `json:"id"`
}

func CreateLink(w http.ResponseWriter, r *http.Request) {

	var payload CreateLinkInput
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := Decode(r.Body, &payload); err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.Info("Creating link for " + payload.Email + " to " + payload.Destination)

	id, err := db.NewLink("", payload.Email, payload.Destination)

	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	output := CreateLinkOutput{ID: id}

	if err := Encode(w, output); err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
