package db

import (
	
	"fmt"

	"github.com/CapregSoft/Hawqal-go/models"
)

/*
   GetCitiesDB() function is supposed to get the cities from the sqlite db
   and return them as a slice of Cities.
*/
func GetCitiesDB(c *Database) ([]*models.Cities, error) {
	//db.Query() returns the rows after selecting attributes from cities table
	rows, err := c.db.Query("SELECT country_id, country_name, state_id, city_id, name FROM cities")
	if err != nil {
		return nil, err
	}

	// defer closes rows after the function executes
	defer rows.Close()

	// creating cities dynamically
	cities := make([]*models.Cities, 0)
	for rows.Next() {
		//interate till last rows
		var city models.Cities

		// scan function scans and returns an err if scan fails
		err := rows.Scan(&city.CountryID, &city.CountryName, &city.StateID, &city.CityID, &city.CityName)
		if err != nil {
			return nil, err
		}
		// appened city into an dynamic cities
		cities = append(cities, &city)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error getting countries from db: %s", err.Error())
	}

	// returns an dynamically created array of cities
	return cities, nil
}
