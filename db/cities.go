package db

import (
	"database/sql"

	"github.com/CapregSoft/Hawqal-go/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

/*
   GetCitiesDB() function is supposed to get the cities from the sqlite db
   and return them as a slice of Cities.
*/
func GetCitiesDB(db *sql.DB, choice *models.Filter) ([]*models.Cities, error) {

	if choice != nil {
		cities := make([]*models.Cities, 0)
		if choice.CityCoordinates {
			if choice.CountryName != "" {
				statePascalCase := cases.Title(language.Und).String(choice.CountryName)
				rows, err := db.Query("SELECT country_name, state_id, city_name, longitude ,latitude FROM cities where country_name=$1", statePascalCase)
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
			} else {
				rows, err := db.Query("SELECT country_name, state_id, city_name, longitude ,latitude FROM cities")
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
			return cities, nil
		}
	}

	cities := make([]*models.Cities, 0)

	if choice != nil {
		if choice.CountryName != "" {
			statePascalCase := cases.Title(language.Und).String(choice.CountryName)
			rows, err := db.Query("SELECT country_name, city_name FROM cities where country_name=$1", statePascalCase)
			if err != nil {
				return nil, err
			}
			defer rows.Close()
			for rows.Next() {
				// interate till last row
				var city models.Cities

				// scan functions scans and returns an err if scan fails
				err := rows.Scan(&city.CountryName, &city.CityName)
				if err != nil {
					return nil, err
				}

				// appened city into an dynamic states
				cities = append(cities, &city)
			}
			return cities, nil
		}
	}

	//db.Query() returns the rows after selecting attributes from cities table
	rows, err := db.Query("SELECT country_name, state_id, city_id, city_name FROM cities")
	if err != nil {
		return nil, err
	}

	// defer closes rows after the function executes
	defer rows.Close()
	for rows.Next() {
		//interate till last rows
		var city models.Cities

		// scan function scans and returns an err if scan fails
		err := rows.Scan(&city.CountryName, &city.StateID, &city.CityID, &city.CityName)
		if err != nil {
			return nil, err
		}
		// appened city into an dynamic cities
		cities = append(cities, &city)
	}
	// returns an dynamically created array of cities
	return cities, nil
}
