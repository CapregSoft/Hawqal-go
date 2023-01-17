package filter

import (
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
		// field :=
		if val.Field(i).Bool() {
			fieldName = append(fieldName, val.Type().Field(i).Name)
		}
	}
	return strings.Join(fieldName, ", ")
}

// func GetTrue(filter interface{}) string {
// 	val := reflect.ValueOf(filter).Elem()
// 	var fieldName []string
// 	for i := 0; i < val.NumField(); i++ {
// 		if val.Field(i).Bool() {
// 			fieldName = append(fieldName, val.Type().Field(i).Name)
// 		}
// 	}
// 	return strings.Join(fieldName, ", ")
// }
