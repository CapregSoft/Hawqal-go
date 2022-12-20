package services

import (
	"database/sql"
	"log"

	db "github.com/CapregSoft/Hawqal-go/db"
	conn "github.com/CapregSoft/Hawqal-go/db_migrator"
	"github.com/CapregSoft/Hawqal-go/models"

	_ "github.com/mattn/go-sqlite3"
)

/*
	The file read.go stimulates the services for the Hawqal-go package
	The functions Gets the data from the sqliteDB and store into the models-structures
*/

//conn.DbConnection() opens the sqlitDB File
var dbConn = conn.DBConnection()

//structure created for db-connection
type DBConnection struct {
	db *sql.DB
}

//the function returns the Database connection
func Database(conn *DBConnection) *sql.DB {
	conn.db = dbConn
	return conn.db
}

//GetCountriesData retreives the data from DB module
//And returns to the example module w.r.t []*models.Country array
func GetCountriesData() ([]*models.Countries, error) {
	countries, err := db.GetCountriesDB(Database(&DBConnection{}))

	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	return countries, nil
}

//GetStatesData retreives the data from DB module GetStatesDB()
//And returns to the example module including an array []*models.States
func GetStatesData() ([]*models.States, error) {
	states, err := db.GetStatesDB(Database(&DBConnection{}))

	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	return states, nil
}

//GetCitiesData retreives the data from DB module GetCitiesDB()
//And returns to the example module with array []*models.Cities
func GetCitiesData() ([]*models.Cities, error) {

	cities, err := db.GetCitiesDB(Database(&DBConnection{}))

	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	//defer executes at when the function executes
	//USer with the function inm order to close the connection to database
	defer Database(&DBConnection{}).Close()

	return cities, nil
}
