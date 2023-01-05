package hawqal

import (
	db "github.com/CapregSoft/Hawqal-go/db"
	"github.com/CapregSoft/Hawqal-go/models"
)

// GetCountriesData retreives the data from Database module
// and return as a []*models.Country.
func GetCountriesData(choice ...*models.CountriesFilter) ([]byte, error) {
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
	defer conn.Close()
	return countriesData, nil
}

// GetStatesData retreives the data from DB module GetStatesDB()
// and return as a []*models.States.
func GetStatesData(choice ...*models.StatesFilter) ([]byte, error) {
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
	defer conn.Close()
	return statesData, nil
}

// GetCitiesData retreives the data from database module GetCitiesDB()
// and return as a slice of Cities.
func GetCitiesData(choice ...*models.CitiesFilter) ([]byte, error) {
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
	defer conn.Close()
	return citiesData, nil
}
