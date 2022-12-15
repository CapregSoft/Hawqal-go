package db

import (
	"database/sql"
	"log"

	"github.com/CapregSoft/Hawqal-go/models"
)

func GetCountriesDB(db *sql.DB) ([]*models.Countries, error) {
	rows, err := db.Query("SELECT country_id, country_name FROM countries ORDER BY country_id ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	countries := make([]*models.Countries, 0)
	for rows.Next() {
		var country models.Countries
		err := rows.Scan(&country.CountryID, &country.CountryName)
		if err != nil {
			return nil, err
		}
		countries = append(countries, &country)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
		return nil, err
	}

	return countries, nil
}
