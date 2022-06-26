package main

import (
	"t02smith.com/url-shortener/db"
)

func main() {
	database := db.OpenConnection()
	db.CreateUrlTable(database)
	defer database.Close()
}
