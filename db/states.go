package db

import (
	"database/sql"
	"log"

	mod "hawqal/models"
)

// GetStates function
func GetStates(db *sql.DB) ([]mod.States, error) {
	rows, err := db.Query("SELECT country_id, country_name, state_id, name FROM states")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	states := []mod.States{}
	for rows.Next() {
		var state mod.States
		err := rows.Scan(&state.CountryID, &state.CountryName, &state.StateID, &state.StateName)
		if err != nil {
			return nil, err
		}
		states = append(states, state)
	}
	
	err = rows.Err()
	if err != nil {
		// Handle the error.
		log.Fatalf("Error executing query: %v", err)
		return nil, err
	}
	return states, nil
}
