package services

import (
    "database/sql"
    "log"
	mod "hawqal/models"
)

// GetCountries function
func GetCountries(db *sql.DB) ([]mod.Countries, error) {
    rows, err := db.Query("SELECT country_id, country_name FROM countries ORDER BY country_id ASC")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    countries := []mod.Countries{}
    for rows.Next() {
        var country mod.Countries
        err := rows.Scan(&country.CountryID, &country.CountryName)
        if err != nil {
            return nil, err
        }
        countries = append(countries, country)
    }

    err = rows.Err()
    if err != nil {
        // Handle the error.
        log.Fatalf("Error executing query: %v", err)
        return nil, err
    }

    return countries, nil
}
