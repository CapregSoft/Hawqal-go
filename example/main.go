package main

import (
	"fmt"
	"log"

	hawqal "github.com/CapregSoft/Hawqal-go"
	"github.com/CapregSoft/Hawqal-go/models"
)

func main() {

	CountriesWithAttributtes, err := hawqal.GetCountriesData(&models.CountryFilter{CountryName: "pakistan"})
	if err != nil {
		log.Fatalf("%v", err)
	}
	for _, country := range CountriesWithAttributtes {
		fmt.Print(string(country))
	}

	states, err := hawqal.GetStatesData(&models.StateFilter{CountryName: "pakistan", StatesCoordinates: true})
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Print("\nStates\n")
	for _, state := range states {
		fmt.Print(string(state))
	}

	cities, err := hawqal.GetCitiesData(&models.CityFilter{CountryName: "pakistan", StateName: "sindh", CityCoordinates: true})
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Print("Cities")
	for _, city := range cities {
		fmt.Print(string(city))
	}
}
