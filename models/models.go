package models

/*
The package models defines the structure of countries,states,cities.
Which is supposed to store the countries,cities,states data.
*/
type Countries struct {
	CountryID   *int    `db:"country_id"`
	CountryName *string `db:"country_name"`
}

type States struct {
	CountryID   *int    `db:"country_id"`
	CountryName *string `db:"country_name"`
	StateID     *int    `db:"state_id"`
	StateName   *string `db:"name"`
}

type Cities struct {
	CountryID   *int    `db:"country_id"`
	CountryName *string `db:"country_name"`
	StateID     *int    `db:"state_id"`
	CityID      *int    `db:"city_id"`
	CityName    *string `db:"name"`
}
