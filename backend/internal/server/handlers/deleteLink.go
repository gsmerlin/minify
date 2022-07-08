package handlers

import (
	"net/http"

	"github.com/gsmerlin/minify/backend/internal/db"
	"github.com/gsmerlin/minify/backend/internal/logger"
)

type DeleteLinkOutput struct {
	ID string `json:"id"`
}

func DeleteLink(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	logger.Info("Deleting link for " + id)

	id, err := db.DeleteLink(id)

	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	output := DeleteLinkOutput{ID: id}

	if err := Encode(w, output); err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
