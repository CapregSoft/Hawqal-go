package hawqal

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/CapregSoft/Hawqal-go/db"
	filter "github.com/CapregSoft/Hawqal-go/filterbool"
	"github.com/CapregSoft/Hawqal-go/models"
)

// Function to Check Error
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetCountriesData retreives the data from Database module
// and return as a []byte after marshling into a json format
func GetCountries(choice ...*models.CountryFilter) ([]byte, error) {
	var trueFields string
	countries := make([]*models.Countries, 0)
	if choice != nil {
		userInput := filter.GetValue(choice[0])
		trueFields = filter.GetTrue(userInput)
	}
	if trueFields != "" {
		userInput := filter.GetValue(choice[0])
		trueFields = filter.GetTrue(userInput)
	} else {
		trueFields = "*"
	}
	sqlQuery := "SELECT "
	sqlQuery = sqlQuery + trueFields + " FROM countries "
	// if country_name != "" {
	// 	country_name = strings.Title(country_name)
	// 	sqlQuery = fmt.Sprintf("%v WHERE country_name = '%v'", sqlQuery, country_name)
	// }
	conn, _ := db.DBConnection()
	defer conn.Close()
	fmt.Print(sqlQuery)
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
	coordinatesJson, err := json.MarshalIndent(countries, "", "  ")
	CheckError(err)
	return coordinatesJson, nil
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

func GetStates(country_name string, choice ...*models.StateFilter) ([]byte, error) {
	var trueFields string
	states := make([]*models.States, 0)
	if choice != nil {
		userInput := filter.GetValue(choice[0])
		trueFields = filter.GetTrue(userInput)
	}
	if trueFields != "" {
		userInput := filter.GetValue(choice[0])
		trueFields = filter.GetTrue(userInput)
	} else {
		trueFields = "*"
	}
	sqlQuery := "SELECT "
	sqlQuery = sqlQuery + trueFields + " FROM states "
	if country_name != "" {
		country_name = strings.Title((country_name))
		sqlQuery = sqlQuery + fmt.Sprintf("WHERE country_name = '%v'", country_name)
	}
	conn, _ := db.DBConnection()
	defer conn.Close()
	fmt.Print(sqlQuery)
	rows, err := conn.Queryx(sqlQuery)
	CheckError(err)
	for rows.Next() {
		var state models.States
		if err := rows.StructScan(&state); err != nil {
			fmt.Println(err)
		}
		states = append(states, &state)
	}
	defer rows.Close()
	coordinatesJson, err := json.MarshalIndent(states, "", "  ")
	CheckError(err)
	return coordinatesJson, nil
}

func GetState(country_name, state_name string, choice ...*models.StateFilter) ([]byte, error) {
	var trueFields string
	states := make([]*models.States, 0)
	if choice != nil {
		userInput := filter.GetValue(choice[0])
		trueFields = filter.GetTrue(userInput)
	}
	if trueFields != "" {
		userInput := filter.GetValue(choice[0])
		trueFields = filter.GetTrue(userInput)
	} else {
		trueFields = "*"
	}
	sqlQuery := "SELECT "
	sqlQuery = sqlQuery + trueFields + " FROM states "
	if country_name != "" {
		country_name = strings.Title((country_name))
		state_name = strings.Title((state_name))
		sqlQuery = sqlQuery + fmt.Sprintf("WHERE country_name = '%v' AND state_name = '%v'", country_name, state_name)
	}
	conn, _ := db.DBConnection()
	defer conn.Close()
	fmt.Print(sqlQuery)
	rows, err := conn.Queryx(sqlQuery)
	CheckError(err)
	for rows.Next() {
		var state models.States
		if err := rows.StructScan(&state); err != nil {
			fmt.Println(err)
		}
		states = append(states, &state)
	}
	defer rows.Close()
	coordinatesJson, err := json.MarshalIndent(states, "", "  ")
	CheckError(err)
	return coordinatesJson, nil
}
