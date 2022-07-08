package handlers

import (
	"fmt"
	"net/http"

	"github.com/gsmerlin/minify/internal/db"
	"github.com/gsmerlin/minify/internal/logger"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	link := r.RequestURI[1:]
	logger.Info(fmt.Sprintf("Received request for: %v", link))
	logger.Info("Checking if record exists...")
	records := db.GetLink(link, "")
	if len(records) == 0 {
		logger.Error("Record not found.")
		w.Write([]byte("404 Not Found"))
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
