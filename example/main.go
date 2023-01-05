package main

import (
	"fmt"
	"log"

	hawqal "github.com/CapregSoft/Hawqal-go"
	"github.com/CapregSoft/Hawqal-go/models"
)

func main() {

	// countries varable is an array of type []*models.Countries
	// contains all the country-city_name & country-id
	//Get the extra attributes of the countries by passing filter as a
	CountriesWithAttributtes, err := hawqal.GetCountriesData(&models.Filter{CountryName: "pakistan"})
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	for _, country := range CountriesWithAttributtes {
		fmt.Print(string(country))
	}
	// states varable is an array of type []*models.States
	// contains data of country-states & states-id
	statesCorrdinates := true

	states, err := hawqal.GetStatesData(&models.Filter{CountryName: "pakistan", StatesCoordinates: statesCorrdinates})
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Print("\nStates\n")
	for _, state := range states {
		fmt.Print(string(state))
	}

	cityCoordinates := false
	//cities varable is an array of type []*models.Cities
	//contains data of country(states-cities) & country-id
	cities, err := hawqal.GetCitiesData(&models.Filter{CountryName: "pakistan", CityCoordinates: cityCoordinates})
	if err != nil {
		//validating the error
		log.Fatalf("Error %v", err)
		//if error occurs prompt the error
	}

	fmt.Print("Cities")
	//loop iteration through the []*models.Cities array.
	if cityCoordinates {
		for _, city := range cities {
			//the variable state holds the data of specific state
			//Follwed by the state name & its specific Id - country city_name
			fmt.Printf("City :: %v -  Longitude:  %v   Latitude:  %v \n", *city.CityName, *city.CitiesLongitude, *city.CitiesLatitude)
		}
	} else {
		for _, city := range cities {
			//the variable state holds the data of specific state
			//Follwed by the state name & its specific Id - country city_name
			fmt.Printf("City :: %v   \n", *city.CityName)
		}
	}

	citiesCountryName := "Pakistan"

	citiesByCountry, err := hawqal.GetCitiesByCountryData(citiesCountryName)
	if err != nil {
		//validating the error
		//if error occurs prompt the error
		log.Fatalf("Error %v", err)
	}
	fmt.Print("Cities For Country ", citiesCountryName)
	//loop iteration through the []*models.Cities array.
	for _, city := range citiesByCountry {
		//the variable city holds the data of specific city *models.Cities
		//Follwed by the city city_name & its City Id - country city_name
		fmt.Printf("City :: %v \n", *city.CityName)
	}

	//city country city_name in order to serach for states
	stateName := "Balochistan"

	citiesByState, err := hawqal.GetCitiesByState(stateName)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	fmt.Print("Cities For State ", stateName)
	//loop iteration through the []*models.Cities array.
	for _, city := range citiesByState {
		//the variable city holds the data of specific city *models.Cities
		//Follwed by the city city_name & its City Id - country city_name
		fmt.Printf("City :: %v \n", *city.CityName)
	}
}
