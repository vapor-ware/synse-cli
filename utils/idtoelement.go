package utils

import (
	"reflect"
	"strconv"
)

/*
Given a Rack ID, will return the index of that rack in the
`scanResponse` struct. WARNING: This may not be consistent between
scans!
*/
func RackIDtoElem(rack_id int) int {
	scanResponse, scanerr := UtilScanOnly() // Add error reporting
	if scanerr != nil {
		return 0
	}
	rackidstring := strconv.Itoa(rack_id)
	scanResponsePtr := reflect.ValueOf(&scanResponse.Racks)
	scanResponseValuePtr := scanResponsePtr.Elem()
	for i := 0; i < scanResponseValuePtr.Len(); i++ {
		if scanResponse.Racks[i].RackID == rackidstring {
			return i
		}
	}
	return 0
}

/*
Given a Board ID, will return the *first encountered* index of that
board in the `scanResponse` struct. WARNING: This may not be
consistent between scans!

This is a temporary measure until UUID's are implemented.
*/
func BoardIDtoElem(board_id int) int {
	scanResponse, scanerr := UtilScanOnly() // Add error reporting
	if scanerr != nil {
		return 0
	}
	boardidstring := strconv.Itoa(board_id)
	scanResponsePtr := reflect.ValueOf(&scanResponse.Racks)
	scanResponseValuePtr := scanResponsePtr.Elem()
	for i := 0; i < scanResponseValuePtr.Len(); i++ {
		boardsPtr := reflect.ValueOf(&scanResponse.Racks[i].Boards)
		boardsValuePtr := boardsPtr.Elem()
		for j := 0; j < boardsValuePtr.Len(); j++ {
			if scanResponse.Racks[i].Boards[j].BoardID == boardidstring {
				return j
			}
		}
	}
	return 0
}

/*
Given a Device ID, will return the *first encountered* index of that
device in the `scanResponse` struct. WARNING: This may not be
consistent between scans!

This is a temporary measure until UUID's are implemented.
*/
func DeviceIDtoElem(device_id int) int {
	scanResponse, scanerr := UtilScanOnly() // Add error reporting
	if scanerr != nil {
		return 0
	}
	deviceidstring := strconv.Itoa(device_id)
	scanResponsePtr := reflect.ValueOf(&scanResponse.Racks)
	scanResponseValuePtr := scanResponsePtr.Elem()
	for i := 0; i < scanResponseValuePtr.Len(); i++ {
		boardsPtr := reflect.ValueOf(&scanResponse.Racks[i].Boards)
		boardsValuePtr := boardsPtr.Elem()
		for j := 0; j < boardsValuePtr.Len(); j++ {
			devicePtr := reflect.ValueOf(&scanResponse.Racks[i].Boards[j].Devices)
			deviceValuePtr := devicePtr.Elem()
			for k := 0; k < deviceValuePtr.Len(); k++ {
				if scanResponse.Racks[i].Boards[j].Devices[k].DeviceID == deviceidstring {
					return k
				}
			}
		}
	}
	return 0
}
