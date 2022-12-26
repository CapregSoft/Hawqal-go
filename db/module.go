package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

/*
   DBConnection is supposed to connect to the sqliteDB and return a *sql.DB object.

   DBConnection is further used to fetch the data from database.
*/
func DBConnection() (*sql.DB, error) {
	databaseFile := "./hawqalDB.sqlite"
	if _, err := os.Stat(databaseFile); os.IsNotExist(err) {
		return nil, fmt.Errorf("database file does not exist")
	}

	// connection to the database.
	db, err := sql.Open("sqlite3", databaseFile)
	if err != nil {
		// handle the error.
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("%v", err)
	}
	// returns the DB Connection
	return db, nil
}
