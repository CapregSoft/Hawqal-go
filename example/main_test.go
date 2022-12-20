package main

import (
	"testing"

	"github.com/CapregSoft/Hawqal-go/models"
	"github.com/stretchr/testify/require"
)

/*
	The File main_test.go is supposed to test the main function execution.
	In this file test cases are described for different scenarious
	Uses testing and testify/require libray for test cases.
*/

//TestMainFunction is defined using testing.T as paramater
//Tests the actual and expected result from functions.
func TestMainFunction(t *testing.T) {

	//defined a structure for testcase & initalized it
	TestCases := []struct {
		Name string

		//Variable ActualLength holds the length of function
		ActualLength int

		//ExpectedLength holds the result of the function
		ExpectedLength int

		//The bool varable stats equal or !equal
		ExpectedResult bool
	}{

		//inialized with the actual and expected value of functions
		//defined for countries-states & cities
		{
			ActualLength:   models.TotalCountries,
			ExpectedLength: 250,
			Name:           "countries",
			ExpectedResult: true,
		},
		{
			ActualLength:   models.TotalStates,
			ExpectedLength: 4989,
			Name:           "states",
			ExpectedResult: true,
		},
		{
			ActualLength:   models.TotalCities,
			ExpectedLength: 150710,
			Name:           "cities",
			ExpectedResult: true,
		},
	}

	//Iterates through the testcases
	for _, tcase := range TestCases {

		//the variable i validates the test cases for countries - states & cities
		i := 0
		t.Run(tcase.Name, func(t *testing.T) {
			if i == 1 {
				//iterates when i = 1 for states

				//func require.Equal comapres two interfaces and returns failure or pass
				require.Equal(t, tcase.ExpectedLength, tcase.ActualLength, tcase.ExpectedResult)
			} else if i == 2 {
				//iterates when i = 2 for cities

				//func require.Equal comapres two interfaces and returns failure or pass
				require.Equal(t, tcase.ExpectedLength, tcase.ActualLength, tcase.ExpectedResult)
			}

			//iterates when i = 0 for countries
			//func require.Equal comapres two interfaces and returns failure or pass
			require.Equal(t, tcase.ExpectedLength, tcase.ActualLength, tcase.ExpectedResult)
		})
		//increaments after each iteration
		i++
	}
}
