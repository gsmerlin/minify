package handlers

import (
	"fmt"
	"net/http"

	"github.com/gsmerlin/minify/backend/internal/db"
	"github.com/gsmerlin/minify/backend/internal/logger"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	link := r.RequestURI[1:]
	logger.Info(fmt.Sprintf("Received request for: %v", link))
	logger.Info("Checking if record exists...")
	records, err := db.GetLink(link, "", "")
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	record := records[0]
	logger.Success("Record found!")

	logger.Info("Attempting to add analytics...")

	if err := db.AddAnalytics(link); err != nil {
		return
	}

	logger.Success("Successfully added analytics!")

	http.Redirect(w, r, record.Destination, http.StatusTemporaryRedirect)
}
