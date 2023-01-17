package models

/*
    The package models defines the structure of countries,states,cities,option.
	Which is supposed to store the countries,cities,states data.
	Option struct created for filter purpose.
*/

//Country Structure Added to store the all the country data
type Countries struct {
	Country_name    string `json:",omitempty" db:"country_name"`
	Currency        string `json:",omitempty" db:"currency"`
	Currency_name   string `json:",omitempty" db:"currency_name"`
	Currency_symbol string `json:",omitempty" db:"currency_symbol"`
	Capital         string `json:",omitempty" db:"capital"`
	Iso_code        string `json:",omitempty" db:"iso_code"`
	Phone_code      string `json:",omitempty" db:"phone_code"`
	Region          string `json:",omitempty" db:"region"`
	Subregion       string `json:",omitempty" db:"subregion"`
	Country_domain  string `json:",omitempty" db:"country_domain"`
	Timezone        string `json:",omitempty" db:"timezone"`
	Zone_city       string `json:",omitempty" db:"zone_city"`
	UTC             string `json:",omitempty" db:"UTC"`
	Latitude        string `json:",omitempty" db:"latitude"`
	Longitude       string `json:",omitempty" db:"longitude"`
}

//States Structure created to store city_name and its ID w.r.t country city_name
type States struct {
	Country_name string `json:",omitempty" db:"country_name"`
	State_id     int    `json:",omitempty" db:"state_id"`
	State_name   string `json:",omitempty" db:"state_name"`
	Latitude     string `json:",omitempty" db:"latitude"`
	Longitude    string `json:",omitempty" db:"longitude"`
}

//Struct city created to store city city_name & id w.r.t country city_name & state ID
type Cities struct {
	Country_name string `json:",omitempty" db:"country_name"`
	State_id     int    `json:",omitempty" db:"state_id"`
	State_name   string `json:",omitempty" db:"state_name"`
	City_id      int    `json:",omitempty" db:"city_id"`
	City_name    string `json:",omitempty" db:"city_name"`
	Latitude     string `json:",omitempty" db:"latitude"`
	Longitude    string `json:",omitempty" db:"longitude"`
}

/*
	Structure option created for advanced functionallities in Hawqal Package

	Used option struct as a filter in three functions...
*/
type CountryFilter struct {
	Country_name string
	CountryMeta  *CountryInner
}

type CountryInner struct {
	Currency        bool
	Currency_name   bool
	Currency_symbol bool
	Capital         bool
	Iso_code        bool
	Phone_code      bool
	Region          bool
	Subregion       bool
	Country_domain  bool
	Timezone        bool
	Zone_city       bool
	UTC             bool
	Latitude        bool
	Longitude       bool
}

//Declaring Constants For Test Cases
//Of Countries-States & Cities
const (
	TotalCountries = 250
	TotalCities    = 150710
	TotalStates    = 4989
)

type StateFilter struct {
	Country_name string
	State_id     bool
	State_name   string
	City_id      bool
	City_name    bool
	Latitude     bool
	Longitude    bool
}

type CityFilter struct {
	Country_name string
	State_id     bool
	State_name   string
	City_id      bool
	City_name    bool
	Latitude     bool
	Longitude    bool
}
