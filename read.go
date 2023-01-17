package hawqal

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/CapregSoft/Hawqal-go/db"
	filter "github.com/CapregSoft/Hawqal-go/filters"
	"github.com/CapregSoft/Hawqal-go/models"
)

// GetCountriesData retreives the data from Database module and return as a []byte after marshling into a json format
func GetCountries(choice ...*models.CountryFilter) ([]byte, error) {
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

func GetCountry(country_name string, choice ...*models.CountryFilter) ([]byte, error) {
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
	if country_name != "" {
		country_name = strings.Title(country_name)
		sqlQuery = sqlQuery + fmt.Sprintf(" WHERE country_name = '%v'", country_name)
	} else {
		return nil, fmt.Errorf("Country_name must be set")
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

func GetStates(choice ...*models.StateFilter) ([]byte, error) {
	var trueFields string
	states := make([]*models.States, 0)
	if choice != nil && choice[0].Filter != nil {
		userInput := filter.GetValue(choice[0].Filter)
		if trueFields = filter.GetTrue(userInput); trueFields == "" {
			trueFields = "*"
		}
	} else {
		trueFields = "*"
	}
	sqlQuery := "SELECT "
	sqlQuery = sqlQuery + trueFields + " FROM states "
	if choice != nil && choice[0].CountryName != "" {
		country_name := strings.Title(choice[0].CountryName)
		sqlQuery = sqlQuery + fmt.Sprintf(" WHERE country_name = '%v'", country_name)
	}
	conn, _ := db.DBConnection()
	defer conn.Close()
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
	StateJson, err := json.MarshalIndent(states, "", "  ")
	CheckError(err)
	return StateJson, nil
}

func GetState(state_name string, choice ...*models.StateFilter) ([]byte, error) {
	var trueFields string
	states := make([]*models.States, 0)
	if choice != nil && choice[0].Filter != nil {
		userInput := filter.GetValue(choice[0].Filter)
		if trueFields = filter.GetTrue(userInput); trueFields == "" {
			trueFields = "*"
		}
	} else {
		trueFields = "*"
	}
	sqlQuery := "SELECT "
	sqlQuery = sqlQuery + trueFields + " FROM states "
	if state_name != "" && choice[0].CountryName != "" {
		country_name := strings.Title(choice[0].CountryName)
		state_name = strings.Title(strings.Title(state_name))
		sqlQuery = sqlQuery + fmt.Sprintf(" WHERE country_name = '%v' AND state_name = '%v'", country_name, state_name)
	} else if state_name != "" {
		state_name := strings.Title(state_name)
		sqlQuery = sqlQuery + fmt.Sprintf(" WHERE state_name = '%v'", state_name)
	} else {
		return nil, fmt.Errorf("state_name must be set")
	}
	conn, _ := db.DBConnection()
	defer conn.Close()
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
	StateJson, err := json.MarshalIndent(states, "", "  ")
	CheckError(err)
	return StateJson, nil
}

func GetCities(choice ...*models.CityFilter) ([]byte, error) {
	var trueFields string
	cities := make([]*models.Cities, 0)
	if choice != nil {
		if choice[0].Filter != nil {
			userInput := filter.GetValue(choice[0].Filter)
			if trueFields = filter.GetTrue(userInput); trueFields == "" {
				trueFields = "*"
			}
		} else {
			trueFields = "*"
		}
	} else {
		trueFields = "*"
	}
	sqlQuery := "SELECT "
	sqlQuery = sqlQuery + trueFields + " FROM cities "
	if choice != nil {
		if choice[0].CountryName != "" && choice[0].StateName != "" {
			country_name := strings.Title(choice[0].CountryName)
			state_name := strings.Title(choice[0].StateName)
			sqlQuery = sqlQuery + fmt.Sprintf(" WHERE country_name = '%v' AND state_name = '%v'", country_name, state_name)
		} else if choice[0].CountryName != "" {
			country_name := strings.Title(choice[0].CountryName)
			sqlQuery = sqlQuery + fmt.Sprintf(" WHERE country_name = '%v'", country_name)
		} else if choice[0].StateName != "" {
			state_name := strings.Title(choice[0].StateName)
			sqlQuery = sqlQuery + fmt.Sprintf(" WHERE state_name = '%v'", state_name)
		}
	}
	conn, _ := db.DBConnection()
	defer conn.Close()
	rows, err := conn.Queryx(sqlQuery)
	CheckError(err)
	for rows.Next() {
		var city models.Cities
		if err := rows.StructScan(&city); err != nil {
			fmt.Println(err)
		}
		cities = append(cities, &city)
	}
	defer rows.Close()
	cityJson, err := json.MarshalIndent(cities, "", "  ")
	CheckError(err)
	return cityJson, nil
}

func GetCity(CityName string, choice ...*models.CityFilter) ([]byte, error) {
	var trueFields string
	cities := make([]*models.Cities, 0)
	if choice != nil && choice[0].Filter != nil {
		userInput := filter.GetValue(choice[0].Filter)
		if trueFields = filter.GetTrue(userInput); trueFields == "" {
			trueFields = "*"
		}
	} else {
		trueFields = "*"
	}
	sqlQuery := "SELECT "
	sqlQuery = sqlQuery + trueFields + " FROM cities "

	if CityName != "" {
		CityName = strings.Title(CityName)
		sqlQuery = sqlQuery + fmt.Sprintf("WHERE city_name = '%v'", CityName)
	} else {
		return nil, fmt.Errorf("CityName must be set")
	}

	if choice != nil {
		if choice[0].CountryName != "" && choice[0].StateName != "" {
			country_name := strings.Title(choice[0].CountryName)
			state_name := strings.Title(choice[0].StateName)
			sqlQuery = sqlQuery + fmt.Sprintf(" AND country_name = '%v' AND state_name = '%v'", country_name, state_name)
		} else if choice[0].CountryName != "" {
			country_name := strings.Title(choice[0].CountryName)
			sqlQuery = sqlQuery + fmt.Sprintf(" AND country_name = '%v'", country_name)
		} else if choice[0].StateName != "" {
			state_name := strings.Title(choice[0].StateName)
			sqlQuery = sqlQuery + fmt.Sprintf(" AND state_name = '%v'", state_name)
		}
	}
	conn, _ := db.DBConnection()
	defer conn.Close()
	rows, err := conn.Queryx(sqlQuery)
	CheckError(err)
	for rows.Next() {
		var city models.Cities
		if err := rows.StructScan(&city); err != nil {
			fmt.Println(err)
		}
		cities = append(cities, &city)
	}
	defer rows.Close()
	cityJson, err := json.MarshalIndent(cities, "", "  ")
	CheckError(err)
	return cityJson, nil
}

// Function to Check Error
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
