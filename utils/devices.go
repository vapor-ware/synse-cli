package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/vapor-ware/vesh/client"

	"github.com/fatih/structs"
)

const Scanpath = "scan"

// scanResponse struct holds the response values from a `/scan` operation.
// While it does not contain the complete set of information available, it does
// contain a complete list of the available assets, including racks and boards.
// The structure mirrors the json struture of response from `/scan` and values
// are assigned to appropriate sub structs.
type scanResponse struct {
	Racks []Rack `json:"racks"`
}

type Rack struct {
	Boards []Board `json:"boards"`
	RackID string  `json:"rack_id"`
}

type Board struct {
	BoardID     string   `json:"board_id"`
	Hostnames   []string `json:"hostnames"`
	IPAddresses []string `json:"ip_addresses"`
	Devices     []Device `json:"devices"`
}

type Device struct {
	DeviceID   string `json:"device_id"`
	DeviceInfo string `json:"device_info"`
	DeviceType string `json:"device_type"`
}

type Result struct {
	Rack
	Board
	Device
}

type FilterFunc struct {
	Result
	FilterFn func(r Result) bool
}

type ResultError struct {
	Result
	Error error
}

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

func GetDevices() (chan Result, error) {
	c := make(chan Result)

	vc := client.New()
	status := &scanResponse{}
	failure := new(client.ErrorResponse)
	resp, err := vc.Sling.New().Get(Scanpath).Receive(status, failure)
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
