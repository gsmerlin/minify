package minify

import (
	"net/http"
)

var internalRepo *Repository = &Repository{}

func handler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		link := req.RequestURI
		record := internalRepo.Read(link[1:], "")[0]
		if record == (Record{}) {
			http.NotFound(res, req)
		}
		internalRepo.addAnalytics(link[1:])
		http.Redirect(res, req, record.Destination, http.StatusPermanentRedirect)
	}
}

func StartServer() {
	internalRepo.InitDB()
	go http.ListenAndServe(":8080", handler())
}

func Repo() *Repository { return internalRepo }
