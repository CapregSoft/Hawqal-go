package db

import (
	"fmt"

	"github.com/CapregSoft/Hawqal-go/models"
)

/*
   GetStatesByCountryDB() function is supposed to get the states from the sqlite db
   and return them as a slice of States.
*/
func GetStatesByCountryDB(c *Database, Country string) ([]*models.States, error) {

	//db.Query() returns the rows after selecting attributes from states table
	rows, err := c.db.Query("SELECT country_id, country_name, state_id, name FROM states where country_name=$1", Country)
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

		// scan function scans and returns an err if scan fails
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
