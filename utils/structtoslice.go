package utils

import (
	//"fmt"
	"reflect"
)

func StructToSingleSlice(str []string) ([]string, error) {
	structPtr := reflect.ValueOf(&str)
	structValuePtr := structPtr.Elem()
	output := make([]string, 0)
	for i := 0; i < structValuePtr.Len(); i++ {
		output = append(output, structValuePtr.Field(i).String())
	}
	return output, nil
}
