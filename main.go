package main

import (
	"fmt"
	db "hawqal/database"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open a connection to the database.
	conn := db.DBConnection()
	defer conn.Close()

	// Execute a query.
	rows, err := conn.Query("SELECT country_id, country_name FROM countries ORDER BY country_id ASC")
	if err != nil {
		// Handle the error.
		log.Fatalf("Error executing query: %v", err)
		return
	}
	defer rows.Close()

	// Loop through the returned rows and print the values.
	for rows.Next() {
		var country_id int
		var country_name string
		err = rows.Scan(&country_id, &country_name)
		if err != nil {
			// Handle the error.
			log.Fatalf("Error executing query: %v", err)
			return
		}
		fmt.Println(country_id, country_name)
	}

	// Check for any errors that may have occurred during the iteration.
	err = rows.Err()
	if err != nil {
		// Handle the error.
		log.Fatalf("Error executing query: %v", err)
		return
	}
}
