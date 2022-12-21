package db

import (
	"database/sql"
	"fmt"

	"github.com/CapregSoft/Hawqal-go/models"
)

/*
   GetCountriesDB is supposed to get the countries from the sqlite db
   and return them as a slice of Countries.
*/
func GetCountriesDB(db *sql.DB) ([]*models.Countries, error) {

	// db.Query returns the rows after selecting attributes from countries table
	rows, err := db.Query("SELECT country_id, country_name FROM countries ORDER BY country_id ASC")
	if err != nil {
		return nil, err
	}
	// defer closes rows after the function executes
	defer rows.Close()

	// creating countries dynamically
	countries := make([]*models.Countries, 0)

	for rows.Next() {
		// interate till last rows
		var country models.Countries

		// scan function scans and returns an err if scan fails
		err := rows.Scan(&country.CountryID, &country.CountryName)
		if err != nil {
			return nil, err
		}
		// appened city into an dynamic countries
		countries = append(countries, &country)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error getting countries from db: %s", err.Error())
	}
	// returns an dynamically created array of countries
	return countries, nil
}
