package api

import (
	"encoding/json"
	"log"
	"net/http"

	"t02smith.com/url-shortener/db"
)

const PATH_START string = "/api/v1"

type Request struct {
	Url string
}

func HandleRequests() {
	log.Println("Starting server on port 8080")
	http.HandleFunc(PATH_START+"/getURL", GetURL)
	http.ListenAndServe(":8080", nil)
}

func GetURL(w http.ResponseWriter, r *http.Request) {

	var req Request

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	new_url := db.FetchURL(db.Database, req.Url)
	w.Write([]byte(new_url))
}
