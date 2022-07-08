package handlers

import (
	"net/http"

	"github.com/gsmerlin/minify/internal/db"
	"github.com/gsmerlin/minify/internal/logger"
)

type EditLinkInput struct {
	ID          string `json:"id"`
	Destination string `json:"destination"`
	Email       string `json:"email"`
}

type EditLinkOutput struct {
	ID string `json:"id"`
}

func EditLink(w http.ResponseWriter, r *http.Request) {

	var payload EditLinkInput

	if err := Decode(r.Body, &payload); err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.Info("Editing link for " + payload.Email + " to " + payload.Destination)

	id, err := db.UpdateLink(payload.ID, payload.Email, payload.Destination)

	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	output := EditLinkOutput{ID: id}

	if err := Encode(w, output); err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
