package server

import (
	"log"
	"net/http"

	"github.com/gsmerlin/minify/internal/db"
	"github.com/gsmerlin/minify/internal/utils"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	link := r.RequestURI[1:]
	log.Printf("Received request for: %v\n", link)
	log.Println("Checking if record exists...")
	record := db.GetLink(link, "")[0]
	if record == (db.Record{}) {
		log.Println("Record not found.")
		w.Write([]byte("404 Not Found"))
		return
	}
	log.Println("Record found!")

	log.Println("Attempting to add analytics...")
	if err := db.AddAnalytics(link); err != nil {
		log.Println("Error adding analytics: ", err)
	}

	http.Redirect(w, r, record.Destination, http.StatusTemporaryRedirect)
}

func routes() {
	http.HandleFunc("/", RedirectHandler)
}

func StartServer(logger *utils.Logger) error {
	if logger != nil {
		log.SetOutput(logger)
	}
	db.InitDB()
	routes()
	log.Println("Running server on port 3000")
	return http.ListenAndServe(":3000", nil)
}
