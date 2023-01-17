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

	data, err := hawqal.GetStates(&models.StateFilter{CountryName: "pakistan", Filter: &models.StateData{StateName: true}})
	CheckError(err)
	fmt.Print(string(data))
}
