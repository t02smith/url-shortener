package util

const (

	// location of SQLite3 database
	DATABASE_LOCATION string = "./urls.db"

	// Domain name for shortened URLs
	DOMAIN string = "http://localhost:8080"

	// Size of the shortened url hash
	// 0 < HASH_SIZE <= 40
	HASH_SIZE int = 5

	// Prefix of API call path
	API_PATH string = "/api/v1"

	// Default port to listen on
	PORT string = ":8080"

	// How long a link is valid for -> 7 days
	LINK_LIFETIME int64 = 60 * 60 * 24 * 7
)
