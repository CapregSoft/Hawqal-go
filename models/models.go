package models

/*
    The package models defines the structure of countries,states,cities.
	Which is supposed to store the countries,cities,states data.
*/

//Country Structure Added to store its specific id and country city_name
type Countries struct {
	CountryName    *string `db:"country_name"`
	Currency       *string `db:"currency"`
	CurrencyName   *string `db:"currency_name"`
	CurrencySymbol *string `db:"currency_symbol"`
	Capital        *string `db:"capital"`
	ISOCode        *string `db:"iso_code"`
	PhoneCode      *string `db:"phone_code"`
	Region         *string `db:"region"`
	SubRegion      *string `db:"subregion"`
	CountryDomain  *string `db:"country_domain"`
	CountryZone    *string `db:"timezone"`
	ZoneCity       *string `db:"zone_city"`
	UTC            *string `db:"UTC"`
	Latitude       *string `db:"latitude"`
	Longitude      *string `db:"longitude"`
}

//States Structure created to store city_name and its ID w.r.t country city_name
type States struct {
	CountryName     *string `db:"country_name"`
	StateID         *int    `db:"state_id"`
	StateName       *string `db:"name"`
	StatesLatitude  *string `db:"latitude"`
	StatesLongitude *string `db:"longitude"`
}

//Struct city created to store city city_name & id w.r.t country city_name & state ID
type Cities struct {
	CountryName     *string `db:"country_name"`
	StateID         *int    `db:"state_id"`
	CityID          *int    `db:"city_id"`
	CityName        *string `db:"city_name"`
	CitiesLatitude  *string `db:"latitude"`
	CitiesLongitude *string `db:"longitude"`
}

type Option struct {
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
