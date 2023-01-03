package db

import (
	"database/sql"

	"github.com/CapregSoft/Hawqal-go/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

/*
   GetStatesDB function is supposed to get the states from the sqlite db and
   return them as a slice of States.
*/
func GetStatesDB(db *sql.DB, choice *models.Option) ([]*models.States, error) {
	states := make([]*models.States, 0)
	if choice != nil {
		if choice.StatesCoordinates {
			if choice.CountryName != "" {
				statePascalCase := cases.Title(language.Und).String(choice.CountryName)
				rows, err := db.Query("SELECT country_name, state_name ,longitude ,latitude FROM states where country_name=$1", statePascalCase)
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
					err := rows.Scan(&state.CountryName, &state.StateName, &state.StatesLongitude, &state.StatesLatitude)
					if err != nil {
						return nil, err
					}

					// appened city into an dynamic states
					states = append(states, &state)
				}
				return states, nil
			} else {
				rows, err := db.Query("SELECT country_name, state_name ,longitude ,latitude FROM states")
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
					err := rows.Scan(&state.CountryName, &state.StateName, &state.StatesLongitude, &state.StatesLatitude)
					if err != nil {
						return nil, err
					}

					// appened city into an dynamic states
					states = append(states, &state)
				}
				return states, nil
			}
		}
	}
	if choice != nil {
		if choice.CountryName != "" {
			statePascalCase := cases.Title(language.Und).String(choice.CountryName)
			rows, err := db.Query("SELECT country_name, state_name FROM states where country_name=$1", statePascalCase)
			if err != nil {
				return nil, err
			}
			defer rows.Close()

			// creating states dynamically
			// states := make([]*models.States, 0)

			for rows.Next() {
				// interate till last row
				var state models.States

				// scan functions scans and returns an err if scan fails
				err := rows.Scan(&state.CountryName, &state.StateName)
				if err != nil {
					return nil, err
				}

				// appened city into an dynamic states
				states = append(states, &state)
			}
			return states, nil
		}
	}
	//db.Query() returns the rows after selecting attributes from states table
	rows, err := db.Query("SELECT country_name, state_id, state_name FROM states")
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
		err := rows.Scan(&state.CountryName, &state.StateID, &state.StateName)
		if err != nil {
			return nil, err
		}

		// appened city into an dynamic states
		states = append(states, &state)
	}
	return states, nil
}
