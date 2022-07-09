package handlers

import (
	"fmt"
	"net/http"

	"github.com/gsmerlin/minify/backend/internal/db"
	"github.com/gsmerlin/minify/backend/internal/logger"
)

type GetAnalyticsOutput struct {
	ID         string   `json:"id"`
	Timestamps []string `json:"timestamps"`
}

func GetAnalytics(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	link := r.URL.Query().Get("id")
	logger.Info(fmt.Sprintf("Received request for: %v", link))
	logger.Info("Checking if record exists...")
	details, err := db.GetAnalytics(link)
	if err != nil {
		logger.Error(err.Error())
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	output := GetAnalyticsOutput{
		ID: details.ID,
	}

	for _, a := range details.Analytics {
		output.Timestamps = append(output.Timestamps, a.Timestamp)
	}

	if err := Encode(w, output); err != nil {
		logger.Error(err.Error())
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}
}
