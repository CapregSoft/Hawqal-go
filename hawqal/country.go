package hawqal

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/CapregSoft/Hawqal-go/db"
	filter "github.com/CapregSoft/Hawqal-go/filters"
	"github.com/CapregSoft/Hawqal-go/models"
)

func GetCountries(choice ...*models.CountryFilter) ([]byte, error) {
	var trueFields string
	if choice != nil {
		userInput := filter.GetValue(choice[0])
		if trueFields = filter.GetTrue(userInput); trueFields == "" {
			trueFields = "*"
		}
	} else {
		trueFields = "*"
	}
	sqlQuery := "SELECT "
	sqlQuery = sqlQuery + trueFields + " FROM countries "
	countries := make([]*models.Countries, 0)
	conn, err := db.DBConnection()
	filter.CheckError(err)
	defer conn.Close()
	rows, err := conn.Queryx(sqlQuery)
	filter.CheckError(err)
	for rows.Next() {
		var country models.Countries
		if err := rows.StructScan(&country); err != nil {
			fmt.Println(err)
		}
		countries = append(countries, &country)
	}
	defer rows.Close()
	CountryJson, err := json.MarshalIndent(countries, "", "  ")
	filter.CheckError(err)
	return CountryJson, nil
}

func GetCountry(CountryName string, choice ...*models.CountryFilter) ([]byte, error) {
	var trueFields string
	countries := make([]*models.Countries, 0)
	if choice != nil {
		userInput := filter.GetValue(choice[0])
		if trueFields = filter.GetTrue(userInput); trueFields == "" {
			trueFields = "*"
		}
	} else {
		trueFields = "*"
	}

	sqlQuery := "SELECT "
	sqlQuery = sqlQuery + trueFields + " FROM countries "
	if CountryName != "" {
		CountryName = strings.Title(CountryName)
		sqlQuery = sqlQuery + fmt.Sprintf(" WHERE country_name = '%v'", CountryName)
	} else {
		return nil, fmt.Errorf("Country_name must be set")
	}
	conn, err := db.DBConnection()
	filter.CheckError(err)
	defer conn.Close()
	rows, err := conn.Queryx(sqlQuery)
	filter.CheckError(err)
	for rows.Next() {
		var country models.Countries
		if err := rows.StructScan(&country); err != nil {
			fmt.Println(err)
		}
		countries = append(countries, &country)
	}
	defer rows.Close()
	CountryJson, err := json.MarshalIndent(countries, "", "  ")
	filter.CheckError(err)
	return CountryJson, nil
}
