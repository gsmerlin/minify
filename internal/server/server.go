package server

import (
	"net/http"

	"github.com/gsmerlin/minify/internal/logger"
	"github.com/gsmerlin/minify/internal/server/handlers"
)

func favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../favicon.ico")
}

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create", handlers.CreateLink)
	mux.HandleFunc("/favicon.ico", favicon)
	mux.HandleFunc("/", handlers.Redirect)
	return mux
}

func Start() error {
	mux := routes()
	logger.Info("Running server on port 3000")
	return http.ListenAndServe(":3000", mux)
}
