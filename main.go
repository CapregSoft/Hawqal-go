package main

import (
	"fmt"
	"log"

	service "github.com/CapregSoft/Hawqal-go/services"
)

func main() {
	// ! Get countries.
	countries, err := service.GetCountriesData()
	if err != nil {
		log.Fatalf("Error %v", err)
	} else {
		fmt.Print("Country")
		for _, country := range countries {
			fmt.Printf("ID :: %v - Country :: %v\n", *country.CountryID, *country.CountryName)
		}
	}
	// ! Get states.
	states, err := service.GetStatesData()
	if err != nil {
		log.Fatalf("Error %v", err)
	} else {
		fmt.Print("States")
		for _, state := range states {
			fmt.Printf("State :: %v - Country :: %v \n", *state.StateName, *state.CountryName)
		}
	}
	cities, err := service.GetCitiesData()
	if err != nil {
		log.Fatalf("Error %v", err)
	} else {
		fmt.Print("Cities")
		for _, city := range cities {
			fmt.Printf("City :: %v - Country :: %v \n", *city.CityName, *city.CountryName)
		}
	}
}
