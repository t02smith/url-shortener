package api

import (
	"encoding/json"
	"net/http"

	"t02smith.com/url-shortener/db"
)

// Replies with a shortened url or an error code
func GetURL(w http.ResponseWriter, r *http.Request) {

	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || !CheckURL(req.Url) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error generating url"))
		return
	}

	new_url := db.FetchURL(db.Database, req.Url)
	w.Write([]byte(new_url))
}
