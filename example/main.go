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
	data, err := hawqal.GetCountries(&models.CountryFilter{Country_name: "pakistan", CountryMeta: &models.CountryInner{Currency_name: false}})
	CheckError(err)
	fmt.Print(string(data))

}
