package db

import (
	"database/sql"
	"log"
	"os"

	"t02smith.com/url-shortener/util"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB = OpenConnection()

// Opens a new database connection
func OpenConnection() *sql.DB {
	_, error := os.Stat(util.DATABASE_LOCATION)

	if error != nil {
		if !os.IsNotExist(error) {
			log.Fatal(error)
		} else {
			f, err := os.Create(util.DATABASE_LOCATION)
			if err != nil {
				log.Fatal(err)
			}
			f.Close()
		}
	}

	database, err := sql.Open("sqlite3", util.DATABASE_LOCATION)

	if err != nil {
		log.Fatal(err)
	}

	CreateUrlTable(database)
	return database
}

// Write a new record to the url table
func WriteUrl(database *sql.DB, row *DatabaseRow) {
	statement, err := database.Prepare(`INSERT INTO urls VALUES (
			?, ?, ?, ?
		);`)

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Inserting %s, %s, %s, %d\n", row.hash, row.old_link, row.New_link, row.expiry)
	_, err = statement.Exec(row.hash, row.old_link, row.New_link, row.expiry)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%s -> %s: Record added successfully\n", row.old_link, util.DOMAIN+row.New_link)
}
