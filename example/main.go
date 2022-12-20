package main

import (
	"fmt"
	"log"

	//imported hawqal module in order to fetch the functions
	hawqal "github.com/CapregSoft/Hawqal-go"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

/*
	This file is suppposed to run the example test for the project.
	Where it will get the data from the hawqal module
	Which retreives the data from (sourse) sqlite Database.
*/

func main() {

	//	countries varable is an array of type []*models.Countries
	//contains all the country-name & country-id
	countries, err := hawqal.GetCountriesData()
	if err != nil {
		//validating the error
		log.Fatalf("Error %v", err)
	}
	fmt.Print("Country")

	//Iterate through the countries array.
	for _, country := range countries {
		//the variable country holds the one country name &  its specific Id
		fmt.Printf("ID :: %v - Country :: %v\n", *country.CountryID, *country.CountryName)
	}

	//states varable is an array of type []*models.States
	//contains data of country-states & states-id
	states, err := hawqal.GetStatesData()
	if err != nil {
		//validating the error
		//if error occurs prompt the error
		log.Fatalf("Error %v", err)
	}

	fmt.Print("States")
	//loop iteration through the []*models.States array.
	for _, state := range states {
		//the variable state holds the data of specific state
		//Follwed by the state name & its specific Id - country name
		fmt.Printf("State :: %v - Country :: %v \n", *state.StateName, *state.CountryName)
	}

	//cities varable is an array of type []*models.Cities
	//contains data of country(states-cities) & country-id
	cities, err := hawqal.GetCitiesData()
	if err != nil {
		//validating the error
		log.Fatalf("Error %v", err)
		//if error occurs prompt the error
	}

	fmt.Print("Cities")
	//loop iteration through the []*models.Cities array.
	for _, city := range cities {
		//the variable city holds the data of specific city *models.Cities
		//Follwed by the city name & its City Id - country name
		fmt.Printf("City :: %v - Country :: %v \n", *city.CityName, *city.CountryName)
	}

	//country name in order to serach for states
	countryName := "Pakistan"

	//The pkg cases &b language used to convert the first letter to upper case
	toUpperCase := cases.Title(language.Und).String(countryName)
	statesByCountry, err := hawqal.GetStatesByCountry(toUpperCase)
	if err != nil {
		//validating the error
		//if error occurs prompt the error
		log.Fatalf("Error %v", err)
	}

	if toUpperCase != "" {
		fmt.Print("\nStates For Country :: ", toUpperCase)
		//loop iteration through the []*models.States array.
		for _, state := range statesByCountry {
			//the variable state holds the data of specific state
			//Follwed by the state name & its specific Id - country name
			fmt.Printf("\nState :: %v  ", *state.StateName)
		}
	}

	//city country name in order to serach for states
	citiesCountryName := "Pakistan"

	//The pkg cases & language used to convert the first letter to upper case
	countryNameToUpper := cases.Title(language.Und).String(citiesCountryName)
	citiesByCountry, err := hawqal.GetCitiesByCountryData(countryNameToUpper)
	if err != nil {
		//validating the error
		//if error occurs prompt the error
		log.Fatalf("Error %v", err)
	}

	if countryNameToUpper != "" {
		fmt.Print("\nCities For Country :: ", countryNameToUpper)
		//loop iteration through the []*models.States array.
		for _, city := range citiesByCountry {
			//the variable state holds the data of specific state
			//Follwed by the state name & its specific Id - country name
			fmt.Printf("\nCities :: %v  ", *city.CityName)
		}
	}
}
