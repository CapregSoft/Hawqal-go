package services

import (
	"database/sql"
	"log"

	db "github.com/CapregSoft/Hawqal-go/db"
	conn "github.com/CapregSoft/Hawqal-go/db_migrator"
	"github.com/CapregSoft/Hawqal-go/models"

	_ "github.com/mattn/go-sqlite3"
)

var dbConn = conn.DBConnection()

type DBConnection struct {
	db *sql.DB
}

func Database(conn *DBConnection) *sql.DB {
	conn.db = dbConn
	return conn.db
}

func GetCountriesData() ([]*models.Countries, error) {
	// Get countries.
	countries, err := db.GetCountriesDB(Database(&DBConnection{}))

	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	return countries, nil
}

// ! Get states function
func GetStatesData() ([]*models.States, error) {

	// Get states.
	states, err := db.GetStatesDB(Database(&DBConnection{}))

	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	return states, nil
}

// ! Get cities function
func GetCitiesData() ([]*models.Cities, error) {

	//  Get cities.
	cities, err := db.GetCitiesDB(Database(&DBConnection{}))

	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}

	defer Database(&DBConnection{}).Close()
	return cities, nil
}
