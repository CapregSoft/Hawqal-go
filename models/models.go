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
	CountryName    bool `db:"country_name"`
	Currency       bool `db:"currency"`
	CurrencyName   bool `db:"currency_name"`
	CurrencySymbol bool `db:"currency_symbol"`
	Capital        bool `db:"capital"`
	IsoCode        bool `db:"iso_code"`
	PhoneCode      bool `db:"phone_code"`
	Region         bool `db:"region"`
	Subregion      bool `db:"subregion"`
	CountryDomain  bool `db:"country_domain"`
	Timezone       bool `db:"timezone"`
	Zone_city      bool `db:"zone_city"`
	UTC            bool `db:"UTC"`
	Latitude       bool `db:"latitude"`
	Longitude      bool `db:"longitide"`
}

//Declaring Constants For Test Cases
//Of Countries-States & Cities
const (
	TotalCountries = 250
	TotalCities    = 150710
	TotalStates    = 4989
)

type StateFilter struct {
	CountryName string `db:"country_name"`
	Filter      *StateData
}

type StateData struct {
	CountryName bool `db:"country_name"`
	StateName   bool `db:"state_name"`
	StateId     bool `db:"state_id"`
	Latitude    bool `db:"latitude"`
	Longitude   bool `db:"longitude"`
}

type CityFilter struct {
	CountryName string `db:"country_name"`
	StateName   string `db:"state_name"`
	Filter      *CityData
}

type CityData struct {
	CountryName bool `db:"country_name"`
	StateName   bool `db:"state_name"`
	StateId     bool `db:"state_id"`
	CityId      bool `db:"city_id"`
	CityName    bool `db:"city_name"`
	Latitude    bool `db:"latitude"`
	Longitude   bool `db:"longitude"`
}
