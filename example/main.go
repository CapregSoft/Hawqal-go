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
	data, err := hawqal.GetCountry("", &models.CountryFilter{Currency: true, Country_name: true})
	CheckError(err)
	fmt.Print(string(data))
	// countryModel := &models.CountryFilter{Currency: true,
	// 	Currency_name: true, Currency_symbol: true, Capital: true, Iso_code: true,
	// 	Phone_code: true, Region: true, Subregion: true, Country_domain: true,
	// 	Time_zone: true, Zonecity: true, UTC: true, Latitude: true, Longitude: true}

	// output := filter.GetTrue(countryModel)
	// fmt.Print(output)
	// CountriesWithAttributtes, err := hawqal.GetCountriesData(&models.CountryFilter{})
	// if err != nil {
	// 	log.Fatalf("%v", err)
	// }
	// for _, country := range CountriesWithAttributtes {
	// 	fmt.Print(string(country))
	// }

	// states, err := hawqal.GetStatesData(&models.StateFilter{CountryName: "pakistan", StatesCoordinates: true})
	// if err != nil {
	// 	log.Fatalf("%v", err)
	// }
	// fmt.Print("\nStates\n")
	// for _, state := range states {
	// 	fmt.Print(string(state))
	// }

	// cities, err := hawqal.GetCitiesData(&models.CityFilter{CountryName: "pakistan", StateName: "sindh", CityCoordinates: true})
	// if err != nil {
	// 	log.Fatalf("%v", err)
	// }
	// fmt.Print("Cities")
	// for _, city := range cities {
	// 	fmt.Print(string(city))
	// }
}
