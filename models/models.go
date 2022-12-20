package models

/*
The package models defines the structure of countries,states,cities.
Which is supposed to store the countries,cities,states data.
*/

//Country Structure Added to store its specific id and country name
type Countries struct {
	CountryID   *int    `db:"country_id"`
	CountryName *string `db:"country_name"`
}

//States Structure created to store name and its ID w.r.t country name
type States struct {
	CountryID   *int    `db:"country_id"`
	CountryName *string `db:"country_name"`
	StateID     *int    `db:"state_id"`
	StateName   *string `db:"name"`
}

//Struct city created to store city name & id w.r.t country name & state ID
type Cities struct {
	CountryID   *int    `db:"country_id"`
	CountryName *string `db:"country_name"`
	StateID     *int    `db:"state_id"`
	CityID      *int    `db:"city_id"`
	CityName    *string `db:"name"`
}

//Declaring Constants For Test Cases
//Of Countries-States & Cities
const (
	TotalCountries = 250
	TotalCities    = 150710
	TotalStates    = 4989
)
