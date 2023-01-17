package main

import (
	"fmt"
	"log"

	hawqal "github.com/CapregSoft/Hawqal-go"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal((err))
	}
}
func main() {

	countries, err := hawqal.GetCountries()
	CheckError(err)
	fmt.Print(string(countries))

	country, err := hawqal.GetCountry("pakistan")
	CheckError(err)
	fmt.Print(string(country))

	states, err := hawqal.GetStates()
	CheckError(err)
	fmt.Print(string(states))

	state, err := hawqal.GetState("Punjab")
	CheckError(err)
	fmt.Print(string(state))

	cities, err := hawqal.GetCities()
	CheckError(err)
	fmt.Print(string(cities))

	city, err := hawqal.GetCity("Lahore")
	CheckError(err)
	fmt.Print(string(city))

}
