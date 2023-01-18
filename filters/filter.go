package filter

import (
	"log"
	"reflect"
	"strings"
)

func GetValue(filter interface{}) interface{} {
	return filter
}

func GetTrue(filter interface{}) string {
	val := reflect.ValueOf(filter).Elem()
	var fieldName []string
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Bool() {
			field := val.Type().Field(i).Tag.Get("db")
			fieldName = append(fieldName, field)
		}
	}
	return strings.Join(fieldName, ", ")
}

// Function to Check Error
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
