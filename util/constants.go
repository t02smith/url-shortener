package util

const (

	// location of SQLite3 database
	DATABASE_LOCATION string = "./urls.db"

	// Domain name for shortened URLs
	DOMAIN string = "http://link.t02smith.com/"

	// Size of the shortened url hash
	// 0 < HASH_SIZE <= 40
	HASH_SIZE int = 5

	// Prefix of API call path
	API_PATH string = "/api/v1"

	// Default port to listen on
	PORT string = ":8080"
)
