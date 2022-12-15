package main

import (
	"fmt"
	"log"

	//imported services module in order to fetch the functions
	service "github.com/CapregSoft/Hawqal-go/services"
)

/*
this file is suppposed to run the test for the project.
where it will get the data from the services module
which retreives the data from (sourse) sqlite Database.
*/

func main() {

	//countries varable is an array of type []*models.Countries
	//contains all the country-name & country-id
	countries, err := service.GetCountriesData()
	if err != nil {
		//validating the error
		log.Fatalf("Error %v", err)
	} else {
		//if err==nil indicates no error returned by GetCountriesData() function
		fmt.Print("Country")

		//Iterate through the countries array.
		for _, country := range countries {
			//the variable country holds the one country name &  its specific Id
			fmt.Printf("ID :: %v - Country :: %v\n", *country.CountryID, *country.CountryName)
		}
	}

	//states varable is an array of type []*models.States
	//contains data of country-states & states-id
	states, err := service.GetStatesData()
	if err != nil {
		//validating the error
		//if error occurs prompt the error
		log.Fatalf("Error %v", err)
	} else {
		//if err==nil indicates no error returned by GetStatesData() function
		fmt.Print("States")

		//loop iteration through the []*models.States array.
		for _, state := range states {
			//the variable state holds the data of specific state
			//Follwed by the state name & its specific Id - country name
			fmt.Printf("State :: %v - Country :: %v \n", *state.StateName, *state.CountryName)
		}
	}

	//cities varable is an array of type []*models.Cities
	//contains data of country(states-cities) & country-id
	cities, err := service.GetCitiesData()
	if err != nil {
		//validating the error
		log.Fatalf("Error %v", err)
		//if error occurs prompt the error
	} else {
		//if err==nil indicates no error returned by GetStatesData() function
		fmt.Print("Cities")

		//loop iteration through the []*models.Cities array.
		for _, city := range cities {
			//the variable city holds the data of specific city *models.Cities
			//Follwed by the city name & its City Id - country name
			fmt.Printf("City :: %v - Country :: %v \n", *city.CityName, *city.CountryName)
		}
	}
}
