package api

type NewURLResponse struct {

	// shortened URL
	shortPath string

	// isoformat date of link expiry
	expiry string
}
