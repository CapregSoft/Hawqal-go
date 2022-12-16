package db

import (
	"database/sql"
	"testing"

	conn "github.com/CapregSoft/Hawqal-go/db_migrator"
)

/*
	Test file for DB is created
	In order to test the database functions
	Test uses the package testing
*/

type Functions interface {
	GetCitiesDB(*sql.DB)
	GetCountriesDB(*sql.DB)
	GetStatesDB(*sql.DB)
}

func TestDB(t *testing.T) {
	db := conn.DBConnection()
	statesDB, err := GetStatesDB(db)
	if statesDB == nil || err != nil {
		t.Errorf("\"db_test.go\" FAILED, expected -> states-data, got -> %v", err)
	}

	countriesDB, err := GetStatesDB(db)
	if countriesDB == nil || err != nil {
		t.Errorf("\"db_test.go\" FAILED, expected -> countries-data, got -> %v", err)
	}

	citiesDB, err := GetStatesDB(db)
	if citiesDB == nil || err != nil {
		t.Errorf("\"db_test.go\" FAILED, expected -> cities-data, got -> %v", err)
	}
}
