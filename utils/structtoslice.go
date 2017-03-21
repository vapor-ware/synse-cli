package utils

import (
	//"fmt"
	"reflect"
)

// StructToSingleSlice takes in a one dimensional struct slice of unknown elements,
// iterates over them, and returns the value of each as a one dimensional string
// slice. Values within the []struct are accessed using the reflect.Elem()
// function.
func StructToSingleSlice(str []string) ([]string, error) {
	structPtr := reflect.ValueOf(&str)
	structValuePtr := structPtr.Elem()
	output := make([]string, 0)
	for i := 0; i < structValuePtr.Len(); i++ {
		output = append(output, structValuePtr.Field(i).String())
	}
	return output, nil // TODO: Add error reporting
}
