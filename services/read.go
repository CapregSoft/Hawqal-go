package services

import (
	"fmt"
	db "hawqal/db"
	dbConn "hawqal/db_migrator"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Get countries function
func GetCountriesData(){

	conn := dbConn.DBConnection()
	defer conn.Close()

	// Get countries.
	countries, err := db.GetCountries(conn)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	fmt.Println("Countries:")
	for _, country := range countries {
		fmt.Printf("%v %v	\n", *country.CountryID, *country.CountryName)
	}
 
}

// Get states function
func GetStatesData() {

	conn := dbConn.DBConnection()
	defer conn.Close()

	// Get states.
	states, err := db.GetStates(conn)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	fmt.Println("States:")
	for _, state := range states {
		fmt.Printf("%v %v %v %v	\n", *state.CountryID, *state.CountryName, *state.StateID, *state.StateName)
	}

}

// Get cities function
func GetCitiesData() {

	conn := dbConn.DBConnection()
	defer conn.Close()

	// Get cities.
	cities, err := db.GetCities(conn)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}

	fmt.Println("Cities:")

	for _, city := range cities {

		fmt.Printf("%v %v %v %v %v	\n", *city.CountryID, *city.CountryName, *city.StateID, *city.CityID, *city.CityName)

	}
}
