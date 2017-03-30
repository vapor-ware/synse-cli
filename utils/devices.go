package utils

import (
	"fmt"
	"net/http"
	"errors"

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
	Path string // FIXME: This shouldn't be here. It should mainstreamed as it is different from DeviceType
	FilterFn func(r Result) bool
}

func FilterDevices(ff *FilterFunc) (chan Result, error) {
	c := make(chan Result)
	errChan := make(chan error)
	fn := ff.FilterFn

	tempchan, err := GetDevices() // FIXME: This should be nested in the function
	if err == nil {
		go func()  {
			var success bool // FIXME: I think there's a better way to do this. Channel length == 0 ?
			for res := range tempchan {
				if fn(res) {
					success = true
					c <- Result{res.Rack, res.Board, res.Device}
				}
			}
			if !success {
				errChan <- DeviceNotFoundErr(ff.Result)
				fmt.Println(errChan)
				panic("balls")
			}

			close(c)
		}()
	}
	err = <-errChan

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
