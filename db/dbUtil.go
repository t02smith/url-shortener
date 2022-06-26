package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DATABASE string = "./urls.db"
)

type DatabaseRow struct {
	hash     string // the sha1 hash of the old domain
	old_link string // the original domain to redirect to
	new_link string // The offset of the hash substring for the link
}

func OpenConnection() *sql.DB {
	database, err := sql.Open("sqlite3", DATABASE)

	if err != nil {
		log.Fatal(err)
	}

	return database
}

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

	log.Println("Record added successfully")
}
