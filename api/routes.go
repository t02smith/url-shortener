package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"t02smith.com/url-shortener/util"
)

// Main routing function to handle requests
func HandleRequests() {
	log.Println("Starting server on port " + util.PORT)

	r := mux.NewRouter()
	r.HandleFunc("/", Index).Methods("GET", "POST")
	r.HandleFunc("/{shortUrl}", RedirectURL).Methods("GET")
	r.HandleFunc(util.API_PATH+"/newURL", NewURL).Methods("POST")

	go http.ListenAndServe(util.PORT, r)

	log.Println("Serving files on port :6060")
	fs := http.NewServeMux()
	fs.Handle("/", http.FileServer(http.Dir("./static/style")))
	http.ListenAndServe(":6060", fs)
}
