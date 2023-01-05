package models

/*
    The package models defines the structure of countries,states,cities,option.
	Which is supposed to store the countries,cities,states data.
	Option struct created for filter purpose.
*/

//Country Structure Added to store the all the country data
type Countries struct {
	CountryName    *string `db:"country_name,omitempty"`
	Currency       *string `db:"currency,omitempty"`
	CurrencyName   *string `db:"currency_name,omitempty"`
	CurrencySymbol *string `db:"currency_symbol,omitempty"`
	Capital        *string `db:"capital,omitempty"`
	ISOCode        *string `db:"iso_code,omitempty"`
	PhoneCode      *string `db:"phone_code,omitempty"`
	Region         *string `db:"region,omitempty"`
	SubRegion      *string `db:"subregion,omitempty"`
	CountryDomain  *string `db:"country_domain,omitempty"`
	CountryZone    *string `db:"timezone,omitempty"`
	ZoneCity       *string `db:"zone_city,omitempty"`
	UTC            *string `db:"UTC,omitempty"`
	Latitude       *string `db:"latitude,omitempty"`
	Longitude      *string `db:"longitude,omitempty"`
}

//States Structure created to store city_name and its ID w.r.t country city_name
type States struct {
	CountryName     *string `db:"country_name,omitempty"`
	StateID         *int    `db:"state_id,omitempty"`
	StateName       *string `db:"name,omitempty"`
	StatesLatitude  *string `db:"latitude,omitempty"`
	StatesLongitude *string `db:"longitude,omitempty"`
}

//Struct city created to store city city_name & id w.r.t country city_name & state ID
type Cities struct {
	CountryName     *string `db:"country_name,omitempty"`
	StateID         *int    `db:"state_id,omitempty"`
	CityID          *int    `db:"city_id,omitempty"`
	CityName        *string `db:"city_name,omitempty"`
	CitiesLatitude  *string `db:"latitude,omitempty"`
	CitiesLongitude *string `db:"longitude,omitempty"`
}

/*
	Structure option created for advanced functionallities in Hawqal Package

	Used option struct as a filter in three functions...
*/
type Filter struct {
	CountryName        string
	Currency           bool
	Region             bool
	Capital            bool
	CountryTime        bool
	CountryCoordinates bool
	CityCoordinates    bool
	StatesCoordinates  bool
}

//Declaring Constants For Test Cases
//Of Countries-States & Cities
const (
	TotalCountries = 250
	TotalCities    = 150710
	TotalStates    = 4989
)
