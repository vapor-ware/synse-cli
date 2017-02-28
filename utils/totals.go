package utils

import (
	"reflect"
)

func TotalElemsNum() int {
	scanResponse, _ := UtilScanOnly() // Add error checking
	var total = 0
	racksPtr := reflect.ValueOf(&scanResponse.Racks)
	racksValuePtr := racksPtr.Elem()
	for i := 0; i < racksValuePtr.Len(); i++ {
		boardsPtr := reflect.ValueOf(&scanResponse.Racks[i].Boards)
		boardsValuePtr := boardsPtr.Elem()
		for j := 0; j < boardsValuePtr.Len(); j++ {
			devicesPtr := reflect.ValueOf(&scanResponse.Racks[i].Boards[j].Devices)
			devicesValuePtr := devicesPtr.Elem()
			for k := 0; k < devicesValuePtr.Len(); k++ {
				total++
			}
		}
	}
	return total
}
