package database

import (
	"database/sql"
	"log"
)

func DBConnection() *sql.DB {
	// Open a connection to the database.
	db, err := sql.Open("sqlite3", "./database/hawqalDB.sqlite")
	if err != nil {
		// Handle the error.
		log.Fatalf("Error opening database: %v", err)
		return nil
	}
	return db
}
