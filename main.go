package main

import (
	dt "hawqal/services"
)

func main() {

	// ! Get countries.
	dt.GetCountriesData()

	// ! Get states.
	dt.GetStatesData()

	// ! Get cities.
	dt.GetCitiesData()
}
