package api

import (
	"log"
	"net/http"
	"time"
)

// Checks a URL to make sure it is a valid link
// \_ a website must return a 200 code from a GET request
func CheckURL(url string) bool {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get("http://" + url)
	if err != nil {
		log.Println(err)
		return false
	}

	return resp.StatusCode >= 200 && resp.StatusCode <= 299
}
