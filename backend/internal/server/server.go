package server

import (
	"net/http"

	"github.com/gsmerlin/minify/backend/internal/logger"
	"github.com/gsmerlin/minify/backend/internal/server/handlers"
)

func favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../favicon.ico")
}

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/analytics", handlers.GetAnalytics)
	mux.HandleFunc("/delete", handlers.DeleteLink)
	mux.HandleFunc("/get", handlers.GetLink)
	mux.HandleFunc("/edit", handlers.EditLink)
	mux.HandleFunc("/create", handlers.CreateLink)
	mux.HandleFunc("/favicon.ico", favicon)
	mux.HandleFunc("/", handlers.Redirect)
	return mux
}

func Start() error {
	mux := routes()
	logger.Info("Running server on port 3001")
	return http.ListenAndServe(":3001", mux)
}
