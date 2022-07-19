package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"t02smith.com/url-shortener/db"
)

// Replies with a shortened url or an error code
func NewURL(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting URL")
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || !CheckURL(req.Url) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error generating url"))
		return
	}

	newUrl := db.RequestURL(db.Database, req.Url, req.Request)

	log.Println("Sending " + newUrl.New_link)

	w.Write([]byte(newUrl.New_link))
}

// Redirect a URL to a given short url if one exists
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl := vars["shortUrl"]
	log.Printf("Redirecting URL %s\n", shortUrl)

	var oldUrl string = db.GetUrlFromShort(db.Database, shortUrl)
	log.Printf("Found %s\n", oldUrl)
	http.Redirect(w, r, "http://"+oldUrl, http.StatusTemporaryRedirect)
	log.Println("Redirected to " + oldUrl)
}
