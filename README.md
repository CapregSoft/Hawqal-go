
# Hawqal-Go
[![GitHub version](https://badge.fury.io/gh/CapregSoft%2FHawqal-go.svg)](https://badge.fury.io/gh/CapregSoft%2FHawqal-go)
[![Godoc](https://godoc.org/github.com/CapregSoft/Hawqal-go?status.svg)](https://godoc.org/github.com/CapregSoft/Hawqal-go)
<!---- [![Build Status](https://travis-ci.com/imgix/imgix-go.svg?branch=main)](https://travis-ci.com/imgix/imgix-go) --->
##
The Golang package that contains the data of world's countries,states and their cities name.This package provide you to get the data of countries, cities & states
## Functions
- Get Countries
- Get States
- Get Cities
- Get Cities By State
- Get State By Country
- Get Cities By Country
## Get Started
```golang
go get github.com/CapregSoft/Hawqal-go
```
## Features/Example
```golang
hawqal.GetCountriesData()   //Returns all the countries & error
hawqal.GetStatesData()     //Returns States & error
hawqal.GetCitiesData()     //Returns all cities & error (if failure occurs)
hawqal.GetStatesByCountry(CountryName)      //Returns States for specific country.
hawqal.GetCitiesByCountryData(CountryName)  //Returns Cities for specific country.
hawqal.GetCitiesByState(StateName)      //Returns Cities for specific State.
```
## Success Response
```bash
- Countries : ['Afghanistan', 'Aland Islands', 'Albania', 'Algeria', . . . ]
- States : ['Alabama', 'Alaska', 'American Samoa', 'Arizona', . . . ]
- Cities : ['Haripur','Abbotabad','WahCantt','Topi',........]
```
## Tech Used
- Golang 1.18
## Authors
- [Capregsoft](https://www.github.com/capregsoft)
- [Ahmed Saleem](https://www.github.com/malikahmed2z2)
