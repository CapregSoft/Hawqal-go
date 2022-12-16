package main

import (
	"testing"

	hawqal "github.com/CapregSoft/Hawqal-go"
	"github.com/CapregSoft/Hawqal-go/models"
)

func TestMainFunction(t *testing.T) {
	type testcases struct {
		Countries []*models.Countries
		States    []*models.States
		Cities    []*models.Cities
	}
	var Test testcases

	countries, countriesErr := hawqal.GetCountriesData()
	Test.Countries = countries
	if countriesErr != nil || countries == nil {
		t.Errorf("\"GetCountriesData()\" FAILED, expected -> countries-data, got -> %v", countriesErr)
	} else {
		t.Logf("\"GetCountriesData()\" SUCCEDED, expected -> countries-data, got -> countries-data")
	}

	states, statesErr := hawqal.GetStatesData()
	Test.States = states
	if statesErr != nil || states == nil {
		t.Errorf("\"GetStatesData()\" FAILED, expected -> states-data, got -> %v", statesErr)
	} else {
		t.Logf("\"GetStatesData()\" SUCCEDED, expected -> states-data, got -> states-data")
	}

	cities, citiesErr := hawqal.GetCitiesData()
	Test.Cities = cities
	if citiesErr != nil || cities == nil {
		t.Errorf("\"GetCitiesData()\" FAILED, expected -> cities-data, got -> %v", citiesErr)
	} else {
		t.Logf("\"GetCitiesData()\" SUCCEDED, expected -> cities-data, got -> cities-data")
	}
}
