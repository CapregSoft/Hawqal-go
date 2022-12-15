package db

import (
	"database/sql"
	"log"

	"github.com/CapregSoft/Hawqal-go/models"
)

// GetStates function
func GetStatesDB(db *sql.DB) ([]*models.States, error) {
	rows, err := db.Query("SELECT country_id, country_name, state_id, name FROM states")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	states := make([]*models.States, 0)
	for rows.Next() {
		var state models.States
		err := rows.Scan(&state.CountryID, &state.CountryName, &state.StateID, &state.StateName)
		if err != nil {
			return nil, err
		}
		states = append(states, &state)
	}

	err = rows.Err()
	if err != nil {
		// Handle the error.
		log.Fatalf("Error executing query: %v", err)
		return nil, err
	}
	return states, nil
}
