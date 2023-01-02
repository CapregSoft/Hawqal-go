package db

import (
	"database/sql"
	"fmt"

	"github.com/CapregSoft/Hawqal-go/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

/*
   GetCountriesDB is supposed to get the countries from the sqlite db
   and return them as a slice of Countries.
*/
func GetCountriesDB(db *sql.DB, choice *models.Option) ([]*models.Countries, error) {

	if choice != nil {
		countries := make([]*models.Countries, 0)
		if choice.Currency {
			if choice.CountryName != "" {
				statePascalCase := cases.Title(language.Und).String(choice.CountryName)
				rows, err := db.Query(`SELECT country_name, currency, currency_name, currency_symbol from countries where country_name=$1`, statePascalCase)
				if err != nil {
					return nil, fmt.Errorf("error %v", err)
				}
				defer rows.Close()

				// creating countries dynamically

				for rows.Next() {
					var country models.Countries

					// interate till last rows
					// scan function scans and returns an err if scan fails
					err := rows.Scan(&country.CountryName, &country.Currency, &country.CurrencyName, &country.CurrencySymbol)
					if err != nil {
						return nil, err
					}
					countries = append(countries, &country)
				}
			} else {
				rows, err := db.Query(`SELECT country_name, currency, currency_name, currency_symbol from countries`)
				if err != nil {
					return nil, fmt.Errorf("error  %v", err)
				}
				defer rows.Close()

				// creating countries dynamically

				for rows.Next() {
					var country models.Countries

					// interate till last rows

					// scan function scans and returns an err if scan fails
					err := rows.Scan(&country.CountryName, &country.Currency, &country.CurrencyName, &country.CurrencySymbol)
					if err != nil {
						return nil, err
					}
					// appened city into an dynamic countries
					countries = append(countries, &country)
				}
			}
		}

		if choice.Region {
			if choice.CountryName != "" {
				statePascalCase := cases.Title(language.Und).String(choice.CountryName)
				rows, err := db.Query(`SELECT country_name, region, subregion, country_domain from countries where country_name=$1 `, statePascalCase)
				if err != nil {
					return nil, fmt.Errorf("error %v", err)
				}
				defer rows.Close()

				// creating countries dynamically

				for rows.Next() {
					var country models.Countries

					// interate till last rows
					// scan function scans and returns an err if scan fails
					err := rows.Scan(&country.CountryName, &country.Region, &country.SubRegion, &country.CountryDomain)
					if err != nil {
						return nil, err
					}
					countries = append(countries, &country)
				}
			} else {
				rows, err := db.Query(`SELECT country_name, region, subregion, country_domain from countries`)
				if err != nil {
					return nil, fmt.Errorf("error  %v", err)
				}
				defer rows.Close()
				for rows.Next() {
					var country models.Countries

					// interate till last rows
					// scan function scans and returns an err if scan fails
					err := rows.Scan(&country.CountryName, &country.Region, &country.SubRegion, &country.CountryDomain)
					if err != nil {
						return nil, err
					}
					// appened city into an dynamic countries
					countries = append(countries, &country)
				}
			}
		}

		if choice.Capital {
			if choice.CountryName != "" {
				statePascalCase := cases.Title(language.Und).String(choice.CountryName)
				rows, err := db.Query(`SELECT country_name ,capital, phone_code, iso_code from countries where country_name=$1 `, statePascalCase)
				if err != nil {
					return nil, fmt.Errorf("im here  %v", err)
				}
				defer rows.Close()
				for rows.Next() {
					var country models.Countries
					// interate till last row
					// scan function scans and returns an err if scan fails
					err := rows.Scan(&country.CountryName, &country.Capital, &country.PhoneCode, &country.ISOCode)
					if err != nil {
						return nil, err
					}
					countries = append(countries, &country)
				}
			} else {
				rows, err := db.Query(`SELECT country_name,capital, phone_code, iso_code from countries`)
				if err != nil {
					return nil, fmt.Errorf("im here  %v", err)
				}
				defer rows.Close()

				// creating countries dynamically
				for rows.Next() {
					var country models.Countries
					// interate till last rows

					// scan function scans and returns an err if scan fails
					err := rows.Scan(&country.CountryName, &country.Capital, &country.PhoneCode, &country.ISOCode)
					if err != nil {
						return nil, err
					}
					// appened city into an dynamic countries
					countries = append(countries, &country)
				}
			}
		}

		if choice.CountryTime {
			if choice.CountryName != " " {
				statePascalCase := cases.Title(language.Und).String(choice.CountryName)
				rows, err := db.Query(`SELECT country_name, timezone ,zone_city, UTC from countries where country_name=$1 `, statePascalCase)
				if err != nil {
					return nil, fmt.Errorf("im here  %v", err)
				}
				defer rows.Close()
				for rows.Next() {
					var country models.Countries
					// interate till last row
					// scan function scans and returns an err if scan fails
					err := rows.Scan(&country.CountryName, &country.CountryZone, &country.ZoneCity, &country.UTC)
					if err != nil {
						return nil, err
					}
					countries = append(countries, &country)
				}
			} else {
				rows, err := db.Query(`SELECT timezone ,zone_city, UTC from countries`)
				if err != nil {
					return nil, fmt.Errorf("im here  %v", err)
				}
				defer rows.Close()

				// creating countries dynamically
				for rows.Next() {
					var country models.Countries
					// interate till last rows

					// scan function scans and returns an err if scan fails
					err := rows.Scan(&country.CountryZone, &country.ZoneCity, &country.UTC)
					if err != nil {
						return nil, err
					}
					// appened city into an dynamic countries
					countries = append(countries, &country)
				}
			}
		}

		if choice.CountryCoordinates {
			if choice.CountryName != " " {
				statePascalCase := cases.Title(language.Und).String(choice.CountryName)
				rows, err := db.Query(`SELECT country_name, longitude ,latitude from countries where country_name=$1 `, statePascalCase)
				if err != nil {
					return nil, fmt.Errorf("im here  %v", err)
				}
				defer rows.Close()
				for rows.Next() {
					var country models.Countries
					// interate till last row
					// scan function scans and returns an err if scan fails
					err := rows.Scan(&country.CountryName, &country.Longitude, &country.Latitude)
					if err != nil {
						return nil, err
					}
					countries = append(countries, &country)
				}
			} else {
				rows, err := db.Query(`SELECT country_name, longitude ,latitude from countries`)
				if err != nil {
					return nil, fmt.Errorf("im here  %v", err)
				}
				defer rows.Close()

				// creating countries dynamically
				for rows.Next() {
					var country models.Countries
					// interate till last rows

					// scan function scans and returns an err if scan fails
					err := rows.Scan(&country.CountryZone, &country.ZoneCity, &country.UTC)
					if err != nil {
						return nil, err
					}
					// appened city into an dynamic countries
					countries = append(countries, &country)
				}
			}
		}
		return countries, nil
	}
	countries := make([]*models.Countries, 0)
	// db.Query returns the rows after selecting attributes from countries table
	rows, err := db.Query("SELECT country_name FROM countries ORDER BY country_name ASC")
	if err != nil {
		return nil, err
	}
	// defer closes rows after the function executes
	defer rows.Close()

	for rows.Next() {
		// interate till last rows
		var country models.Countries

		// scan function scans and returns an err if scan fails
		err := rows.Scan(&country.CountryName)
		if err != nil {
			return nil, err
		}
		// appened city into an dynamic countries
		countries = append(countries, &country)
	}
	// returns an dynamically created array of countries
	return countries, nil
}
