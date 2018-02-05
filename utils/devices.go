package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/vapor-ware/synse-cli/client"

	"github.com/fatih/structs"
)

// Scanpath contains the api path for performing scans
const Scanpath = "scan"

// scanResponse struct holds the response values from a `/scan` operation.
// While it does not contain the complete set of information available, it does
// contain a complete list of the available assets, including racks and boards.
// The structure mirrors the json struture of response from `/scan` and values
// are assigned to appropriate sub structs.
type scanResponse struct {
	Racks []Rack `json:"racks"`
}

// Rack contains the top level objects for a rack
type Rack struct {
	Boards []Board `json:"boards"`
	RackID string  `json:"id"`
}

// Board contains the top level objects for a board
type Board struct {
	BoardID     string   `json:"id"`
	Hostnames   []string `json:"hostnames"`
	IPAddresses []string `json:"ip_addresses"`
	Devices     []Device `json:"devices"`
}

// Device contains the response values for a specific device
type Device struct {
	DeviceID   string `json:"id"`
	DeviceInfo string `json:"info"`
	DeviceType string `json:"type"`
}

// Result gathers the response values for all nested objects
type Result struct {
	Rack
	Board
	Device
}

// FilterFunc contains a matching function for conditions to be satisfied when
// parsing the full list of devices. It should return `true` for any device satisfying
// the search conditions.
type FilterFunc struct {
	Result
	FilterFn func(r Result) bool
}

// ResultError contains response fields for query errors from the API.
type ResultError struct {
	Result
	Error error
}

// FilterDevices takes in a FilterFunc object and parses the full list of devices
// for matches. If there are no errors, it will return a slice of the matching objects.
func FilterDevices(ff *FilterFunc) (chan ResultError, error) {
	c := make(chan ResultError)
	fn := ff.FilterFn

	tempchan, err := GetDevices() // FIXME: This should be nested in the function
	if err == nil {
		go func() {
			var success bool // FIXME: I think there's a better way to do this. Channel length == 0 ?
			for res := range tempchan {
				if fn(res) {
					success = true
					c <- ResultError{Result{res.Rack, res.Board, res.Device}, nil}
				}
			}
			if !success {
				var res Result
				c <- ResultError{res, DeviceNotFoundErr(ff.Result)}
			}

			close(c)
		}()
	}

	return c, err
}

// GetDevices queries the endpoint for a summary of all devices. It then walks
// the tree and populates rack, board, and device responses for each object. The
// resulting Result object can be passed to other functions for filtering.
func GetDevices() (chan Result, error) {
	c := make(chan Result)

	status := &scanResponse{}
	failure := new(client.ErrorResponse)
	resp, err := client.New().Get(Scanpath).Receive(status, failure)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprint(failure))
	}
	fmt.Println("API reported status ok")

	go func() {
		for _, rack := range structs.New(status).Field("Racks").Value().([]Rack) {
			for _, board := range structs.New(rack).Field("Boards").Value().([]Board) {
				for _, device := range structs.New(board).Field("Devices").Value().([]Device) {
					c <- Result{rack, board, device}
				}
			}
		}

		close(c)
	}()

	return c, nil
}
