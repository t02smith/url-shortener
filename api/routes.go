package api

import (
	"log"
	"net/http"
)

const PATH_START string = "/api/v1"

type Request struct {
	Url string
}

// Main routing function to handle requests
func HandleRequests() {
	log.Println("Starting server on port 8080")
	http.HandleFunc(PATH_START+"/getURL", GetURL)
	http.ListenAndServe(":8080", nil)
}
