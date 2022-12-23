package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

/*
   DBConnection is supposed to connect to the sqliteDB and return a *sql.DB object.

   DBConnection is further used to fetch the data from database.
*/
func DBConnection() (*Database, error) {
	// connection to the database.
	db, err := sql.Open("sqlite3", "./HawqalData.sqlite")
	if err != nil {
		// handle the error.
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}
	if err = db.Ping(); err != nil {
		// db.Close()
		return nil, fmt.Errorf("%v", err)
	}
	// returns the DB Connection
	return &Database{db: db}, nil
}
