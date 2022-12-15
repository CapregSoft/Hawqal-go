package db

import (
	"database/sql"
	"log"

	"github.com/CapregSoft/Hawqal-go/models"
)

/*
   GetStatesDB() function is supposed to get the states from the sqlite db.
   This function then return the structure which contains  all the states
*/

func GetStatesDB(db *sql.DB) ([]*models.States, error) {

	//db.Query() returns the rows after selecting attributes from states table

	rows, err := db.Query("SELECT country_id, country_name, state_id, name FROM states")
	if err != nil {
		return nil, err
	}

	// defer closes rows after the function executes
	defer rows.Close()

	//creating states dynamically
	states := make([]*models.States, 0)

	for rows.Next() {
		//interate till last row
		var state models.States

		//Scan() functions scans and returns an err if scan fails
		err := rows.Scan(&state.CountryID, &state.CountryName, &state.StateID, &state.StateName)
		if err != nil {
			return nil, err
		}

		//appened city into an dynamic states
		states = append(states, &state)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
		return nil, err
	}

	//returns an dynamically created array of states
	return states, nil
}
