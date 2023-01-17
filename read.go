package hawqal

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/CapregSoft/Hawqal-go/db"
	filter "github.com/CapregSoft/Hawqal-go/filtering"
	"github.com/CapregSoft/Hawqal-go/models"
)

// GetCountriesData retreives the data from Database module and return as a []byte after marshling into a json format
func GetCountries(choice ...*models.CountryFilter) ([]byte, error) {
	var trueFields string
	countries := make([]*models.Countries, 0)
	if choice != nil && choice[0].CountryMeta != nil {
		userInput := filter.GetValue(choice[0].CountryMeta)
		if trueFields = filter.GetTrue(userInput); trueFields == "" {
			trueFields = "*"
		}
	} else {
		trueFields = "*"
	}
	sqlQuery := "SELECT "
	sqlQuery = sqlQuery + trueFields + " FROM countries "
	if choice != nil && len(choice[0].Country_name) != 0 {
		country_name := strings.Title(choice[0].Country_name)
		sqlQuery = sqlQuery + fmt.Sprintf(" WHERE country_name = '%v'", country_name)
	}
	conn, _ := db.DBConnection()
	defer conn.Close()
	rows, err := conn.Queryx(sqlQuery)
	CheckError(err)
	for rows.Next() {
		var country models.Countries
		if err := rows.StructScan(&country); err != nil {
			fmt.Println(err)
		}
		countries = append(countries, &country)
	}
	defer rows.Close()
	CountryJson, err := json.MarshalIndent(countries, "", "  ")
	CheckError(err)
	return CountryJson, nil
}

// func GetCountry(country_name string, choice ...*models.CountryFilter) ([]byte, error) {
// 	var data string

// 	countries := make([]*models.Countries, 0)
// 	if choice != nil {
// 		output := filter.GetValue(choice[0])
// 		trueFields = filter.GetTrue(output)
// 	}
// 	if data != "" {
// 		output := filter.GetValue(choice[0])
// 		trueFields = filter.GetTrue(output)
// 	} else {
// 		data = "*"
// 	}
// 	sqlQuery := "SELECT "
// 	if country_name != "" {
// 		country_name = strings.Title(country_name)
// 		sqlQuery = fmt.Sprintf("%v%v FROM countries where country_name='%v'", sqlQuery, data, country_name)
// 	} else {
// 		return nil, fmt.Errorf("coutries name is must required ")
// 	}
// 	conn, _ := db.DBConnection()
// 	defer conn.Close()
// 	fmt.Print(sqlQuery)
// 	rows, err := conn.Queryx(sqlQuery)
// 	CheckError(err)
// 	for rows.Next() {
// 		var country models.Countries
// 		if err := rows.StructScan(&country); err != nil {
// 			fmt.Println(err)
// 		}
// 		countries = append(countries, &country)
// 	}
// 	defer rows.Close()
// 	coordinatesJson, err := json.MarshalIndent(countries, "", "  ")
// 	CheckError(err)
// 	return coordinatesJson, nil
// }

// Function to Check Error
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
