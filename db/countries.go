package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/CapregSoft/Hawqal-go/models"
)

/*
   GetCountriesDB is supposed to get the countries from the sqlite db
   and return them as a slice of Countries.
*/
func GetCountriesDB(db *sql.DB, choice *models.CountryFilter) ([]byte, error) {

	if choice != nil {
		countries := make([]*models.Countries, 0)
		if choice.Currency {
			if choice.CountryName != "" {
				choice.CountryName = strings.Title(choice.CountryName)

				rows, err := db.Query(`SELECT country_name, currency, currency_name, currency_symbol from countries where country_name=$1`, choice.CountryName)
				if err != nil {
					return nil, fmt.Errorf("error %v", err)
				}
				defer rows.Close()
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
			currencyJson, err := json.MarshalIndent(countries, "", "  ")
			if err != nil {
				return nil, err
			}
			return currencyJson, nil
		}

		if choice.Region {
			if choice.CountryName != "" {
				choice.CountryName = strings.Title(choice.CountryName)

				rows, err := db.Query(`SELECT country_name, region, subregion, country_domain from countries where country_name=$1 `, choice.CountryName)
				if err != nil {
					return nil, fmt.Errorf("error %v", err)
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
			regionJson, err := json.MarshalIndent(countries, "", "  ")
			if err != nil {
				return nil, err
			}
			return regionJson, nil
		}

		if choice.Capital {
			if choice.CountryName != "" {
				choice.CountryName = strings.Title(choice.CountryName)

				rows, err := db.Query(`SELECT country_name ,capital, phone_code, iso_code from countries where country_name=$1 `, choice.CountryName)
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
			capitalJson, err := json.MarshalIndent(countries, "", "  ")
			if err != nil {
				return nil, err
			}
			return capitalJson, nil
		}

		if choice.CountryTime {
			if choice.CountryName != " " {
				choice.CountryName = strings.Title(choice.CountryName)

				rows, err := db.Query(`SELECT country_name, timezone ,zone_city, UTC from countries where country_name=$1 `, choice.CountryName)
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
			countryTimeJson, err := json.MarshalIndent(countries, "", "  ")
			if err != nil {
				return nil, err
			}
			return countryTimeJson, nil
		}

		if choice.CountryCoordinates {
			if choice.CountryName != " " {
				choice.CountryName = strings.Title(choice.CountryName)

				rows, err := db.Query(`SELECT country_name, longitude ,latitude from countries where country_name=$1 `, choice.CountryName)
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
			coordinatesJson, err := json.MarshalIndent(countries, "", "  ")
			if err != nil {
				return nil, err
			}
			return coordinatesJson, nil
		}

		if choice.CountryName != "" {
			choice.CountryName = strings.Title(choice.CountryName)

			rows, err := db.Query("SELECT * FROM countries WHERE country_name = $1", choice.CountryName)
			if err != nil {
				return nil, err
			}
			// defer closes rows after the function executes
			defer rows.Close()
			for rows.Next() {
				// interate till last rows
				var country models.Countries
				// scan function scans and returns an err if scan fails
				err := rows.Scan(&country.CountryName, &country.ISOCode, &country.PhoneCode, &country.Capital, &country.Currency, &country.CurrencyName, &country.CurrencySymbol, &country.CountryDomain, &country.Region, &country.SubRegion, &country.CountryZone, &country.ZoneCity, &country.UTC, &country.Latitude, &country.Longitude)
				if err != nil {
					return nil, err
				}
				// appened city into an dynamic countries
				countries = append(countries, &country)
			}
			countryJson, err := json.MarshalIndent(countries, "", "  ")
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
			}
			// returns an dynamically created array of countries
			return countryJson, nil
		}
	}
	countries := make([]*models.Countries, 0)
	// db.Query returns the rows after selecting attributes from countries table
	rows, err := db.Query("SELECT * FROM countries ORDER BY country_name ASC")
	if err != nil {
		return nil, err
	}
	// defer closes rows after the function executes
	defer rows.Close()

	for rows.Next() {
		// interate till last rows
		var country models.Countries
		// scan function scans and returns an err if scan fails
		err := rows.Scan(&country.CountryName, &country.ISOCode, &country.PhoneCode, &country.Capital, &country.Currency, &country.CurrencyName, &country.CurrencySymbol, &country.CountryDomain, &country.Region, &country.SubRegion, &country.CountryZone, &country.ZoneCity, &country.UTC, &country.Latitude, &country.Longitude)
		if err != nil {
			return nil, err
		}
		// appened city into an dynamic countries
		countries = append(countries, &country)
	}
	countriesJson, err := json.MarshalIndent(countries, "", "  ")
	if err != nil {
		return nil, err
	}
	// returns an dynamically created array of countries
	return countriesJson, nil
}
