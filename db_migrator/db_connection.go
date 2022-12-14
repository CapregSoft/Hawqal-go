package db_migrator

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

)

func DBConnection() *sql.DB {
	// Open a connection to the database.
	db, err := sql.Open("sqlite3", "./db_migrator/hawqalDB.sqlite")
	if err != nil {
		// Handle the error.
		log.Fatalf("Error opening database: %v", err)
		return nil
	}
	return db
}
