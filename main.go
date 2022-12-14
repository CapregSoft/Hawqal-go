package main

import (
	"fmt"
	db "hawqal/database"
	ser "hawqal/services"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open a connection to the database.

	conn := db.DBConnection()
	defer conn.Close()

	// ! Get countries.
	countries, err := ser.GetCountries(conn)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	fmt.Println("Countries:")
	for _, country := range countries {
		fmt.Printf("%v %v	\n", *country.CountryID, *country.CountryName)
	}

	// ! Get states.
	states, err := ser.GetStates(conn)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	fmt.Println("States:")
	for _, state := range states {
		fmt.Printf("%v %v %v %v	\n", *state.CountryID, *state.CountryName, *state.StateID, *state.StateName)
	}

	// ! Get cities.
	cities, err := ser.GetCities(conn)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}

	fmt.Println("Cities:")

	for _, city := range cities {

		fmt.Printf("%v %v %v %v %v	\n", *city.CountryID, *city.CountryName, *city.StateID, *city.CityID, *city.CityName)

	}

}
