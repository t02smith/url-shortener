package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"t02smith.com/url-shortener/util"
)

// Main routing function to handle requests
func HandleRequests() {
	log.Println("Starting server on port 8080")

	r := mux.NewRouter()
	r.HandleFunc("/{shortUrl}", RedirectURL)
	r.HandleFunc(util.API_PATH+"/getURL", GetURL)

	http.ListenAndServe(util.PORT, r)
}
