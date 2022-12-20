package services

import (
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

//GetCountriesData retreives the data from DB module
//And returns to the example module w.r.t []*models.Country array
func GetCountriesData() ([]*models.Countries, error) {
	//DbConnection() connects to the db in order to retreive the data from it.
	var dbConn = conn.DBConnection()

	//passed db as a paramater in order to connects to the db
	countries, err := db.GetCountriesDB(dbConn)

	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	//Closes the connection at the end when the function executes
	defer dbConn.Close()
	return countries, nil
}

//GetStatesData retreives the data from DB module GetStatesDB()
//And returns to the example module including an array []*models.States
func GetStatesData() ([]*models.States, error) {
	//DbConnection() connects to the db in order to retreive the data from it.
	var dbConn = conn.DBConnection()

	//passed db as a paramater in order to connects to the db
	states, err := db.GetStatesDB(dbConn)

	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	defer dbConn.Close()
	return states, nil
}

//GetCitiesData retreives the data from DB module GetCitiesDB()
//And returns to the example module with array []*models.Cities
func GetCitiesData() ([]*models.Cities, error) {
	//DbConnection() connects to the db in order to retreive the data from it.
	var dbConn = conn.DBConnection()

	//passed db as a paramater in order to connects to the db
	cities, err := db.GetCitiesDB(dbConn)

	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	//defer executes at when the function executes
	//Uses with the function in order to close the connection to database
	// defer Database(&DBConnection{}).Close()
	defer dbConn.Close()
	return cities, nil
}

//GetStatesByCountry accepts the country name as a paramater .
//return the states for the specific country name
func GetStatesByCountry(country string) ([]*models.States, error) {
	//DbConnection() connects to the db in order to retreive the data from it.
	var dbConn = conn.DBConnection()

	//passed *sql.db as a paramater in order to connects to the db
	states, err := db.GetStatesByCountryDB(dbConn, country)

	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	defer dbConn.Close()
	return states, nil
}

//GetCitiesBYCountryData retreives the data from DB module GetCitiesByCountry()
//Accepts the param of string as a country name
//And returns to the example module with array []*models.Cities
func GetCitiesByCountryData(country string) ([]*models.Cities, error) {
	//DbConnection() connects to the db in order to retreive the data from it.
	var dbConn = conn.DBConnection()

	//passed db as a paramater in order to connects to the db

	cities, err := db.GetCitiesByCountry(dbConn, country)

	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	defer dbConn.Close()
	return cities, nil
}

//GetCitiesBYCountryData retreives the data from DB module GetCitiesByCountry()
//Accepts the param of string as a country name
//And returns to the example module with array []*models.Cities
func GetCitiesByState(state string) ([]*models.Cities, error) {
	//DbConnection() connects to the db in order to retreive the data from it.
	var dbConn = conn.DBConnection()

	//passed db as a paramater in order to connects to the db

	cities, err := db.GetCitiesByStateDB(dbConn, state)

	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	//defer executes at when the function executes
	//Uses with the function in order to close the connection to database
	// defer DBConnection().Close()
	defer dbConn.Close()
	return cities, nil
}
