package handlers

import (
	"net/http"

	"github.com/gsmerlin/minify/backend/internal/db"
	"github.com/gsmerlin/minify/backend/internal/logger"
)

type GetLinkOutput struct {
	ID          string `json:"id"`
	Destination string `json:"destination"`
	Email       string `json:"email"`
}

func GetLink(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	logger.Info("Received request for " + email)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	records, err := db.GetLink("", email, "")

	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := Encode(w, records); err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
