package db

import (
	"database/sql"
	"encoding/hex"
	"log"
	"time"

	"t02smith.com/url-shortener/util"
)

// Check if a URL has already been stored in the DB
func UrlExists(database *sql.DB, url string) DatabaseRow {
	hash := hex.EncodeToString(util.Hash(url))
	rows, err := database.Query("SELECT * FROM urls WHERE hash = ?1", hash)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return DatabaseRow{}
	}

	var row DatabaseRow
	rows.Scan(&row.hash, &row.old_link, &row.new_link, &row.expiry)

	// delete record if it has expired
	if (row != DatabaseRow{}) && row.expiry < time.Now().Unix() {
		log.Printf("Link expired for %s. Removing link.\n", url)
		statement, _ := database.Prepare("DELETE FROM urls WHERE hash=?1;")
		statement.Exec(row.hash)

		return DatabaseRow{}
	}

	return row
}

// Check if a url has already been used
func NewUrlExists(database *sql.DB, url string) bool {
	rows, err := database.Query("SELECT * FROM urls WHERE new_link=?1", url)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var exists bool = rows.Next()
	return exists
}

// Fetches the shortened URL from the DB or generates it
func FetchURL(database *sql.DB, url string) string {
	var new_url DatabaseRow = UrlExists(database, url)

	// generate a new URL
	if (DatabaseRow{} == new_url) {
		log.Printf("%s not found. Generating shortened url\n", url)
		var row = GenerateURL(database, url)

		WriteUrl(database, row)
		return util.DOMAIN + row.new_link
	}

	log.Printf("Found %s. Sending %s\n", url, util.DOMAIN+new_url.new_link)
	return util.DOMAIN + new_url.new_link
}

// Generates a new shortened url
func GenerateURL(database *sql.DB, url string) DatabaseRow {
	log.Printf("Generating url for %s...\n", url)

	hash := hex.EncodeToString(util.Hash(url))
	var offset int = 0

	for offset < (len(hash) - util.HASH_SIZE) {
		if NewUrlExists(database, hash[offset:offset+util.HASH_SIZE]) {
			offset++
		} else {
			break
		}
	}

	return DatabaseRow{
		hash:     hash,
		old_link: url,
		new_link: hash[offset : offset+util.HASH_SIZE],
		expiry:   time.Now().Unix() + util.LINK_LIFETIME,
	}
}

// Generates a requested URL
func RequestURL(database *sql.DB, url string, requestedUrl string) string {
	if len(requestedUrl) == 0 {
		return FetchURL(database, url)
	}

	var reqUrl DatabaseRow = UrlExists(database, requestedUrl)

	// url available
	if (reqUrl == DatabaseRow{}) {
		log.Printf("%s not found. Generating shortened url\n", requestedUrl)
		WriteUrl(database, DatabaseRow{
			hash:     hex.EncodeToString(util.Hash(url)),
			old_link: url,
			new_link: requestedUrl,
			expiry:   time.Now().Unix() + util.LINK_LIFETIME,
		})

		return util.DOMAIN + requestedUrl
	}

	log.Printf("%s not available.", requestedUrl)
	return FetchURL(database, url)
}

// Gets a link from the shortened hash value
func GetUrlFromShort(database *sql.DB, shortUrl string) string {
	statement, err := database.Prepare("SELECT * FROM urls WHERE new_link=?")
	if err != nil {
		log.Println(err)
	}

	rows, error := statement.Query(shortUrl)
	if error != nil {
		log.Println(error)
	}

	defer rows.Close()

	var row DatabaseRow
	if rows.Next() {
		rows.Scan(&row.hash, &row.old_link, &row.new_link, &row.expiry)

		if row.expiry >= time.Now().Unix() {
			return row.old_link
		}
	}

	// TODO redirect to not found page
	return util.DOMAIN + "/error/not-found"
}
