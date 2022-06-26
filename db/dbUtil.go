package db

import (
	"crypto/sha1"
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DATABASE_LOCATION string = "./urls.db"
	DOMAIN            string = "http://link.t02smith.com/"
)

var Database *sql.DB = OpenConnection()

type DatabaseRow struct {
	hash     string // the sha1 hash of the old domain
	old_link string // the original domain to redirect to
	new_link string // The offset of the hash substring for the link
}

// Creates a new urls table
func CreateUrlTable(database *sql.DB) {
	_, err := database.Exec(`CREATE TABLE IF NOT EXISTS urls (
			hash TEXT PRIMARY KEY,
			old_link TEXT NOT NULL,
			new_link TEXT NOT NULL
		);`)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("URL table created")
}

// Opens a new database connection
func OpenConnection() *sql.DB {
	f, err := os.Create(DATABASE_LOCATION)
	if err != nil {
		log.Fatal(err)
	}
	f.Close()

	database, err := sql.Open("sqlite3", DATABASE_LOCATION)

	if err != nil {
		log.Fatal(err)
	}

	CreateUrlTable(database)
	return database
}

// Write a new record to the url table
func WriteUrl(database *sql.DB, row DatabaseRow) {
	statement, err := database.Prepare(`INSERT INTO urls VALUES (
			?, ?, ?
		)`)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = statement.Exec(row.hash, row.old_link, row.new_link)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%s -> %s: Record added successfully\n", row.old_link, DOMAIN+row.new_link)
}

// Return the sha1 hash of a string
func Hash(s string) []byte {
	hasher := sha1.New()
	hasher.Write([]byte(s))
	return hasher.Sum(nil)
}
