package db

import (
	"database/sql"
	"log"
	mod "hawqal/models"
)

// GetCities function
func GetCities(db *sql.DB) ([]mod.Cities, error) {
	rows, err := db.Query("SELECT country_id, country_name, state_id, city_id, name FROM cities")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cities := []mod.Cities{}
	for rows.Next() {
		var city mod.Cities
		err := rows.Scan(&city.CountryID, &city.CountryName, &city.StateID, &city.CityID, &city.CityName)
		if err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}

	err = rows.Err()
	if err != nil {
		// Handle the error.
		log.Fatalf("Error executing query: %v", err)
		return nil, err
	}
	return cities, nil
}
