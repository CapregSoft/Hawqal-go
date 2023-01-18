package main

import (
	"fmt"
	"log"

	hawqal "github.com/CapregSoft/Hawqal-go"
	"github.com/CapregSoft/Hawqal-go/models"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal((err))
	}
}
func main() {

	// countries, err := hawqal.GetCountries()
	// CheckError(err)
	// fmt.Print(string(countries))

	// country, err := hawqal.GetCountry("pakistan")
	// CheckError(err)
	// fmt.Print(string(country))

	// states, err := hawqal.GetStates()
	// CheckError(err)
	// fmt.Print(string(states))

	state, err := hawqal.GetState("Punjab", &models.StateFilter{CountryName: "pakistan", Filter: &models.StateData{StateName: true}})
	CheckError(err)
	fmt.Print(string(state))

	// cities, err := hawqal.GetCities()
	// CheckError(err)
	// fmt.Print(string(cities))

	// city, err := hawqal.GetCity("Lahore")
	// CheckError(err)
	// fmt.Print(string(city))

}
