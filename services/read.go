package services

import (
	"database/sql"
	"fmt"
	db "hawqal/db"
	conn "hawqal/db_migrator"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var dbConn = conn.DBConnection()

type DBConnections struct {
	db *sql.DB
}

func database(conn *DBConnections) *sql.DB {
	conn.db = dbConn
	return conn.db
}

// ! Get countries function
func GetCountriesData() {

	// Get countries.
	countries, err := db.GetCountries(database(&DBConnections{}))

	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	fmt.Println("Countries:")
	for _, country := range countries {
		fmt.Printf("%v %v	\n", *country.CountryID, *country.CountryName)
	}

}

// ! Get states function
func GetStatesData() {

	// Get states.
	states, err := db.GetStates(database(&DBConnections{}))

	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	fmt.Println("States:")
	for _, state := range states {
		fmt.Printf("%v %v %v %v	\n", *state.CountryID, *state.CountryName, *state.StateID, *state.StateName)
	}

}

// ! Get cities function
func GetCitiesData() {

	//  Get cities.
	cities, err := db.GetCities(database(&DBConnections{}))

	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}

	fmt.Println("Cities:")

	for _, city := range cities {

		fmt.Printf("%v %v %v %v %v	\n", *city.CountryID, *city.CountryName, *city.StateID, *city.CityID, *city.CityName)

	}

	defer database(&DBConnections{}).Close()
}
