package handlers

import (
	"io"
	"net/http"

	"github.com/gsmerlin/minify/backend/internal/db"
	"github.com/gsmerlin/minify/backend/internal/logger"
)

type CreateLinkInput struct {
	Destination string `json:"destination"`
	Email       string `json:"email"`
}

type CreateLinkOutput struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Destination string `json:"destination"`
}

func CreateLink(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var payload CreateLinkInput

	if err := Decode(r.Body, &payload); err != nil {
		if err == io.EOF {
			return
		}
		logger.Error(err.Error())
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	logger.Info("Creating link for " + payload.Email + " to " + payload.Destination)

	record, err := db.NewLink("", payload.Email, payload.Destination)

	if err != nil {
		logger.Error(err.Error())
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	if err := Encode(w, record); err != nil {
		logger.Error(err.Error())
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}
}
