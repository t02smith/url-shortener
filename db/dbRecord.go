package db

import (
	"database/sql"
	"log"
)

type DatabaseRow struct {
	hash     string // (PRIMARY KEY) the sha1 hash of the old domain
	old_link string // the original domain to redirect to
	new_link string // the offset of the hash substring for the link
	expiry   int64  // timestamp for when it expires
}

// Creates a new urls table
func CreateUrlTable(database *sql.DB) {
	_, err := database.Exec(`CREATE TABLE IF NOT EXISTS urls (
			hash TEXT PRIMARY KEY,
			old_link TEXT NOT NULL,
			new_link TEXT NOT NULL,
			expiry INTEGER NOT NULL
		);`)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("URL table created")
}
