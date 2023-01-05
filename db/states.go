package db

import (
	"database/sql"
	"encoding/json"

	"github.com/CapregSoft/Hawqal-go/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

/*
   GetStatesDB function is supposed to get the states from the sqlite db and
   return them as a slice of States.
*/
func GetStatesDB(db *sql.DB, choice *models.Filter) ([]byte, error) {
	states := make([]*models.States, 0)
	if choice != nil {
		if choice.StatesCoordinates {
			if choice.CountryName != "" {
				statePascalCase := cases.Title(language.Und).String(choice.CountryName)
				rows, err := db.Query("SELECT country_name,state_id, state_name ,longitude ,latitude FROM states where country_name=$1", statePascalCase)
				if err != nil {
					return nil, err
				}

				// defer closes rows after the function executes
				defer rows.Close()

				// creating states dynamically

				for rows.Next() {
					// interate till last row
					var state models.States

					// scan functions scans and returns an err if scan fails
					err := rows.Scan(&state.CountryName, &state.StateID, &state.StateName, &state.StatesLongitude, &state.StatesLatitude)
					if err != nil {
						return nil, err
					}

					// appened city into an dynamic states
					states = append(states, &state)
				}
			} else {
				rows, err := db.Query("SELECT country_name,state_id, state_name ,longitude ,latitude FROM states")
				if err != nil {
					return nil, err
				}

				// defer closes rows after the function executes
				defer rows.Close()

				// creating states dynamically

				for rows.Next() {
					// interate till last row
					var state models.States

					// scan functions scans and returns an err if scan fails
					err := rows.Scan(&state.CountryName, &state.StateID, &state.StateName, &state.StatesLongitude, &state.StatesLatitude)
					if err != nil {
						return nil, err
					}

					// appened city into an dynamic states
					states = append(states, &state)
				}
			}
			coordinatesJson, err := json.MarshalIndent(states, "", "  ")
			if err != nil {
				return nil, err
			}
			return coordinatesJson, nil
		}
		if choice.CountryName != "" {
			statePascalCase := cases.Title(language.Und).String(choice.CountryName)
			rows, err := db.Query("SELECT state_id,state_name,country_name FROM states where country_name=$1", statePascalCase)
			if err != nil {
				return nil, err
			}
			defer rows.Close()
			for rows.Next() {
				// interate till last row
				var state models.States

				// scan functions scans and returns an err if scan fails
				err := rows.Scan(&state.StateID, &state.StateName, &state.CountryName)
				if err != nil {
					return nil, err
				}
				// appened city into an dynamic states
				states = append(states, &state)
			}
			stateJson, err := json.MarshalIndent(states, "", "  ")
			if err != nil {
				return nil, err
			}
			return stateJson, nil
		}
	}

	//db.Query() returns the rows after selecting attributes from states table
	rows, err := db.Query("SELECT state_id,state_name,country_name FROM states ORDER BY state_name ASC")
	if err != nil {
		return nil, err
	}
	// defer closes rows after the function executes
	defer rows.Close()
	for rows.Next() {
		// interate till last row
		var state models.States
		// scan functions scans and returns an err if scan fails
		err := rows.Scan(&state.StateID, &state.StateName, &state.CountryName)
		if err != nil {
			return nil, err
		}
		// appened city into an dynamic states
		states = append(states, &state)
	}
	stateJson, err := json.MarshalIndent(states, "", "  ")
	if err != nil {
		return nil, err
	}
	return stateJson, nil
}
