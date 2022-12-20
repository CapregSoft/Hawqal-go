package db

import (
	"database/sql"
	"fmt"

	"github.com/CapregSoft/Hawqal-go/models"
)

/*
   GetStatesDB function is supposed to get the states from the sqlite db and
   return them as a slice of States.
*/
func GetStatesDB(db *sql.DB) ([]*models.States, error) {

	//db.Query() returns the rows after selecting attributes from states table
	rows, err := db.Query("SELECT country_id, country_name, state_id, name FROM states")
	if err != nil {
		return nil, err
	}

	// defer closes rows after the function executes
	defer rows.Close()

	// creating states dynamically
	states := make([]*models.States, 0)

	for rows.Next() {
		// interate till last row
		var state models.States

		// scan functions scans and returns an err if scan fails
		err := rows.Scan(&state.CountryID, &state.CountryName, &state.StateID, &state.StateName)
		if err != nil {
			return nil, err
		}

		// appened city into an dynamic states
		states = append(states, &state)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error getting countries from db: %s", err.Error())
	}

	// returns an dynamically created array of states
	return states, nil
}
