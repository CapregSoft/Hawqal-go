package db

import (
	"database/sql"
	"log"

	"github.com/CapregSoft/Hawqal-go/models"
)

/*
   GetCitiesByCountry() function is supposed to get the cities by country name from the sqlite db.
   This function then return the structure which contains  all the cities for specific country name
*/

func GetCitiesByStateDB(db *sql.DB, state string) ([]*models.Cities, error) {
	//db.Query() returns the rows after selecting attributes from cities table
	rows, err := db.Query(`SELECT cities.name FROM cities,states WHERE cities.state_id = states.state_id AND States.name =$1`, state)
	if err != nil {
		return nil, err
	}
	// defer closes rows after the function executes
	defer rows.Close()

	//creating cities dynamically
	cities := make([]*models.Cities, 0)
	for rows.Next() {
		//interate till last rows
		var city models.Cities

		//Scan() functions scans and returns an err if scan fails
		err := rows.Scan(&city.CityName)
		if err != nil {
			return nil, err
		}
		//appened city into an dynamic cities
		cities = append(cities, &city)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
		return nil, err
	}
	//returns an dynamically created array of cities
	return cities, nil
}