package hawqal

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/CapregSoft/Hawqal-go/db"
	filter "github.com/CapregSoft/Hawqal-go/filters"
	"github.com/CapregSoft/Hawqal-go/models"
)

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
	conn, err := db.DBConnection()
	filter.CheckError(err)
	defer conn.Close()
	rows, err := conn.Queryx(sqlQuery)
	filter.CheckError(err)
	for rows.Next() {
		var state models.States
		if err := rows.StructScan(&state); err != nil {
			fmt.Println(err)
		}
		states = append(states, &state)
	}

	defer rows.Close()
	StateJson, err := json.MarshalIndent(states, "", "  ")
	filter.CheckError(err)
	return StateJson, nil
}

func GetState(StateName string, choice ...*models.StateFilter) ([]byte, error) {
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
	if choice != nil {
		if StateName != "" && choice[0].CountryName != "" {
			country_name := strings.Title(choice[0].CountryName)
			StateName = strings.Title(strings.Title(StateName))
			sqlQuery = sqlQuery + fmt.Sprintf(" WHERE country_name = '%v' AND state_name = '%v'", country_name, StateName)
		}
	} else if StateName != "" {
		StateName := strings.Title(StateName)
		sqlQuery = sqlQuery + fmt.Sprintf(" WHERE state_name = '%v'", StateName)
	} else {
		return nil, fmt.Errorf("state_name must be set")
	}
	conn, _ := db.DBConnection()
	defer conn.Close()
	// filter.CheckError(err)

	rows, _ := conn.Queryx(sqlQuery)
	for rows.Next() {
		var state models.States
		if err := rows.StructScan(&state); err != nil {
			fmt.Println(err)
		}
		states = append(states, &state)
	}

	defer rows.Close()
	StateJson, err := json.MarshalIndent(states, "", "  ")
	filter.CheckError(err)
	return StateJson, nil
}
