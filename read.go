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
func GetCountriesData(choice ...*models.Filter) ([]byte, error) {
	// connects to the database in order to retreive the data from it.
	conn, err := db.DBConnection()
	if err != nil {
		return nil, err
	}
	// passed db as a paramater in order to connects to the db
	if len(choice) == 0 {
		countries, err := db.GetCountriesDB(conn, nil)
		if err != nil {
			return nil, err
		}
		return countries, nil
	}
	countriesData, err := db.GetCountriesDB(conn, choice[0])
	if err != nil {
		return nil, err
	}
	return countriesData, nil
}

// GetStatesData retreives the data from DB module GetStatesDB()
// and return as a []*models.States.
func GetStatesData(choice ...*models.Filter) ([]*models.States, error) {
	// connects to the database in order to retreive the data from it.
	conn, err := db.DBConnection()
	if err != nil {
		return nil, err
	}

	// passed db as a paramater in order to connects to the db
	if len(choice) == 0 {
		states, err := db.GetStatesDB(conn, nil)
		if err != nil {
			return nil, err
		}
		return states, nil
	}
	statesData, err := db.GetStatesDB(conn, choice[0])
	if err != nil {
		return nil, err
	}
	return statesData, nil
}

// GetCitiesData retreives the data from database module GetCitiesDB()
// and return as a slice of Cities.
func GetCitiesData(choice ...*models.Filter) ([]*models.Cities, error) {
	// connects to the database in order to retreive the data from it.
	conn, err := db.DBConnection()
	if err != nil {
		return nil, err
	}
	if len(choice) == 0 {
		cities, err := db.GetCitiesDB(conn, nil)
		if err != nil {
			return nil, err
		}
		return cities, nil
	}
	citiesData, err := db.GetCitiesDB(conn, choice[0])
	if err != nil {
		return nil, err
	}
	return citiesData, nil
}

// GetStatesByCountry accepts the country city_name as a paramater .
// return the states for the specific country city_name
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
//Accepts the param of string as a country city_name
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
// that accepts the param as a string of country city_name and
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
