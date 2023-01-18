package hawqal

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/CapregSoft/Hawqal-go/db"
	filter "github.com/CapregSoft/Hawqal-go/filters"
	"github.com/CapregSoft/Hawqal-go/models"
)

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
	conn, err := db.DBConnection()
	filter.CheckError(err)
	defer conn.Close()
	rows, err := conn.Queryx(sqlQuery)
	filter.CheckError(err)
	for rows.Next() {
		var city models.Cities
		if err := rows.StructScan(&city); err != nil {
			fmt.Println(err)
		}
		cities = append(cities, &city)
	}
	defer rows.Close()
	cityJson, err := json.MarshalIndent(cities, "", "  ")
	filter.CheckError(err)
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
	conn, err := db.DBConnection()
	filter.CheckError(err)
	defer conn.Close()
	rows, err := conn.Queryx(sqlQuery)
	filter.CheckError(err)
	for rows.Next() {
		var city models.Cities
		if err := rows.StructScan(&city); err != nil {
			fmt.Println(err)
		}
		cities = append(cities, &city)
	}
	defer rows.Close()
	cityJson, err := json.MarshalIndent(cities, "", "  ")
	filter.CheckError(err)
	return cityJson, nil
}
