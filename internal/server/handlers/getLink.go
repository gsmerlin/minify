package handlers

import (
	"net/http"

	"github.com/gsmerlin/minify/internal/db"
	"github.com/gsmerlin/minify/internal/logger"
)

type GetLinkOutput struct {
	ID          string `json:"id"`
	Destination string `json:"destination"`
	Email       string `json:"email"`
}

func GetLink(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	records, err := db.GetLink(id, "")

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
