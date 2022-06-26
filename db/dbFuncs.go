package db

import (
	"database/sql"
	"encoding/hex"
	"log"
)

func UrlExists(database *sql.DB, url string) DatabaseRow {
	hash := hex.EncodeToString(Hash(url))
	rows, err := database.Query("SELECT * FROM urls WHERE hash = ?1;", hash)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return DatabaseRow{}
	}

	var row DatabaseRow
	rows.Scan(&row.hash, &row.old_link, &row.new_link)
	return row

}

func NewUrlExists(database *sql.DB, url string) bool {
	rows, err := database.Query("SELECT * FROM urls WHERE new_link=?1", url)

	if err != nil {
		log.Fatal(err)
	}

	return rows.Next()
}

func FetchURL(database *sql.DB, url string) string {
	var new_url DatabaseRow = UrlExists(database, url)

	if (DatabaseRow{} == new_url) {
		log.Println("No URL found. Generating url")
		var row = GenerateURL(database, url)
		WriteUrl(database, row)
		return row.new_link + DOMAIN
	}

	return new_url.new_link + DOMAIN
}

func GenerateURL(database *sql.DB, url string) DatabaseRow {
	log.Println("Generating url...")
	hash := hex.EncodeToString(Hash(url))
	var offset int = 0

	for offset < (len(hash) - 5) {
		if NewUrlExists(database, hash[offset:offset+5]) {
			offset++
		} else {
			break
		}
	}

	return DatabaseRow{
		hash:     hash,
		old_link: url,
		new_link: hash[offset : offset+5],
	}

}
