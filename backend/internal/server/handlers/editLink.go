package handlers

import (
	"io"
	"net/http"

	"github.com/gsmerlin/minify/backend/internal/db"
	"github.com/gsmerlin/minify/backend/internal/logger"
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
	enableCors(&w)
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	var payload EditLinkInput

	if err := Decode(r.Body, &payload); err != nil {
		if err == io.EOF {
			return
		}
		logger.Error(err.Error())
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	logger.Info("Editing link for " + payload.Email + " to " + payload.Destination)

	id, err := db.UpdateLink(payload.ID, payload.Email, payload.Destination)

	if err != nil {
		logger.Error(err.Error())
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	output := EditLinkOutput{ID: id}

	if err := Encode(w, output); err != nil {
		logger.Error(err.Error())
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}
}
