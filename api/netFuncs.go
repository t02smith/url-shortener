package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"t02smith.com/url-shortener/db"
)

// Replies with a shortened url or an error code
func GetURL(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting URL")
	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || !CheckURL(req.Url) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error generating url"))
		return
	}

	new_url := db.FetchURL(db.Database, req.Url)
	log.Println("Sending " + new_url)
	w.Write([]byte(new_url))
}

// Redirect a URL to a given short url if one exists
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	log.Println("Redirecting URL...")
	vars := mux.Vars(r)
	shortUrl := vars["shortUrl"]

	var oldUrl string = db.GetUrlFromShort(db.Database, shortUrl)
	http.Redirect(w, r, "http://"+oldUrl, http.StatusPermanentRedirect)
	log.Println("Redirected to " + oldUrl)
}
