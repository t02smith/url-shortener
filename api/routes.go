package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PATH_START string = "/api/v1"

type Request struct {
	Url string
}

// Main routing function to handle requests
func HandleRequests() {
	log.Println("Starting server on port 8080")

	r := mux.NewRouter()
	r.HandleFunc("/{shortUrl}", RedirectURL)
	r.HandleFunc(PATH_START+"/getURL", GetURL)

	http.ListenAndServe(":8080", r)
}
