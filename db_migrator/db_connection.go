package db_migrator

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

/*
   The functiuon DBConnection() is supposed to connect to the sqliteDB & opens it
   DbConnection() is further used to fetch the data from Db After opening it.
   Which returns the Db connection.
*/

func DBConnection() *sql.DB {
	// Open a connection to the database.
	db, err := sql.Open("sqlite3", "./db_migrator/hawqalDB.sqlite")
	if err != nil {
		// Handle the error.
		log.Fatalf("Error opening database: %v", err)
		return nil
	}
	//Returns the DB Connection
	return db
}
