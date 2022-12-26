package hawqal

import (
	"fmt"

	db "github.com/CapregSoft/Hawqal-go/db"
	"github.com/CapregSoft/Hawqal-go/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// GetCountriesData retreives the data from Database module
// and return as a []*models.Country.
func GetCountriesData() ([]*models.Countries, error) {
	// connects to the database in order to retreive the data from it.
	conn, err := db.DBConnection()
	if err != nil {
		return nil, err
	}

	// passed db as a paramater in order to connects to the db
	countries, err := db.GetCountriesDB(conn)
	if err != nil {
		return nil, err
	}

	// closes the connection at the end when the function executes
	return countries, nil
}

// GetStatesData retreives the data from DB module GetStatesDB()
// and return as a []*models.States.
func GetStatesData() ([]*models.States, error) {
	// connects to the database in order to retreive the data from it.
	conn, err := db.DBConnection()
	if err != nil {
		return nil, err
	}

	// passed db as a paramater in order to connects to the db
	states, err := db.GetStatesDB(conn)
	if err != nil {
		return nil, err
	}

	// closes the connection at the end when the function executes
	return states, nil
}

// GetCitiesData retreives the data from database module GetCitiesDB()
// and return as a slice of Cities.
func GetCitiesData() ([]*models.Cities, error) {
	// connects to the database in order to retreive the data from it.
	conn, err := db.DBConnection()
	if err != nil {
		return nil, err
	}

	//passed db as a paramater in order to connects to the db
	cities, err := db.GetCitiesDB(conn)
	if err != nil {
		return nil, err
	}

	// defer executes at when the function executes
	// uses with the function in order to close the connection to database
	// defer Database(&DBConnection{}).Close()
	return cities, nil
}

// GetStatesByCountry accepts the country name as a paramater .
// return the states for the specific country name
func GetStatesByCountry(country string) ([]*models.States, error) {
	if len(country) == 0 {
		return nil, fmt.Errorf("country is required")
	}
	// the pkg cases & language used to convert the first letter to pascal case
	statePascalCase := cases.Title(language.Und).String(country)

	// connects to the database in order to retreive the data from it.
	conn, err := db.DBConnection()
	if err != nil {
		return nil, err
	}

	// passed db as a paramater in order to connects to the database.
	states, err := db.GetStatesByCountryDB(conn, statePascalCase)
	if err != nil {
		return nil, err
	}

	return states, nil
}

//GetCitiesBYCountryData retreives the data from DB module GetCitiesByCountry()
//Accepts the param of string as a country name
//And returns to the example module with array []*models.Cities
func GetCitiesByCountryData(country string) ([]*models.Cities, error) {
	if len(country) == 0 {
		return nil, fmt.Errorf("country is required")
	}
	// the pkg cases & language used to convert the first letter to pascal case
	statePascalCase := cases.Title(language.Und).String(country)
	// connects to the database in order to retreive the data from it.
	conn, err := db.DBConnection()
	if err != nil {
		return nil, err
	}

	// passed db as a paramater in order to connects to the database.
	cities, err := db.GetCitiesByCountry(conn, statePascalCase)
	if err != nil {
		return nil, err
	}

	return cities, nil
}

// GetCitiesByState retreives the data from database module GetCitiesByStateDB()
// that accepts the param as a string of country name and
// returns array of []*models.Cities
func GetCitiesByState(state string) ([]*models.Cities, error) {
	if len(state) == 0 {
		return nil, fmt.Errorf("state is required")
	}
	// the pkg cases & language used to convert the first letter to pascal case
	statePascalCase := cases.Title(language.Und).String(state)

	// connects to the database in order to retreive the data from it.
	conn, err := db.DBConnection()
	if err != nil {
		return nil, err
	}

	// passed db as a paramater in order to connects to the db
	cities, err := db.GetCitiesByStateDB(conn, statePascalCase)
	if err != nil {
		return nil, err
	}
	return cities, nil
}
