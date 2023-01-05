package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/CapregSoft/Hawqal-go/models"
)

/*
   GetCitiesDB() function is supposed to get the cities from the sqlite db
   and return them as a slice of Cities.
*/
func GetCitiesDB(db *sql.DB, choice *models.CitiesFilter) ([]byte, error) {

	if choice != nil {
		cities := make([]*models.Cities, 0)
		if choice.CityCoordinates {
			if choice.CountryName != "" {
				choice.CountryName = strings.Title(choice.CountryName)
				rows, err := db.Query("SELECT country_name, state_id, city_name, longitude ,latitude FROM cities where country_name=$1", choice.CountryName)
				if err != nil {
					return nil, err
				}
				// defer closes rows after the function executes
				defer rows.Close()
				for rows.Next() {
					//interate till last rows
					var city models.Cities

					// scan function scans and returns an err if scan fails
					err := rows.Scan(&city.CountryName, &city.StateID, &city.CityName, &city.CitiesLongitude, &city.CitiesLatitude)
					if err != nil {
						return nil, err
					}
					// appened city into an dynamic cities
					cities = append(cities, &city)
				}
			}
			if choice.CountryName != "" && choice.StateName != "" {
				country := strings.Title(choice.CountryName)
				state := strings.Title(choice.StateName)
				rows, err := db.Query("SELECT country_name,city_id,state_id, state_name,city_name, longitude ,latitude FROM cities where country_name=$1 and state_name=$2", country, state)
				if err != nil {
					return nil, err
				}
				// defer closes rows after the function executes
				defer rows.Close()
				for rows.Next() {
					//interate till last rows
					var city models.Cities

					// scan function scans and returns an err if scan fails
					err := rows.Scan(&city.CountryName, &city.CityID, &city.StateID, &city.StateName, &city.CityName, &city.CitiesLongitude, &city.CitiesLatitude)
					if err != nil {
						return nil, err
					}
					// appened city into an dynamic cities
					cities = append(cities, &city)
				}
			}
			coordinatesJson, err := json.MarshalIndent(cities, "", "  ")
			if err != nil {
				return nil, err
			}
			return coordinatesJson, nil
		}

		if choice.CountryName != "" {
			choice.CountryName = strings.Title(choice.CountryName)
			if choice.StateName != "" {
				choice.StateName = strings.Title(choice.StateName)
				rows, err := db.Query("SELECT city_id,city_name,state_name,state_id,country_name FROM cities where country_name=$1  and state_name =$2", choice.CountryName, choice.StateName)
				if err != nil {
					return nil, err
				}
				defer rows.Close()
				for rows.Next() {
					var city models.Cities
					err := rows.Scan(&city.CityID, &city.CityName, &city.StateName, &city.StateID, &city.CountryName)
					if err != nil {
						return nil, err
					}
					cities = append(cities, &city)
				}
			} else {
				rows, err := db.Query("SELECT city_id,city_name,state_name,state_id,country_name FROM cities where country_name=$1", choice.CountryName)
				if err != nil {
					return nil, err
				}
				defer rows.Close()
				for rows.Next() {
					// interate till last row
					var city models.Cities

					// scan functions scans and returns an err if scan fails
					err := rows.Scan(&city.CityID, &city.CityName, &city.StateName, &city.StateID, &city.CountryName)
					if err != nil {
						return nil, err
					}
					// appened city into an dynamic states
					cities = append(cities, &city)
				}
			}
			coordinatesJson, err := json.MarshalIndent(cities, "", "  ")
			if err != nil {
				return nil, err
			}
			return coordinatesJson, nil
		}
	}
	return nil, fmt.Errorf("country name is required")
}
