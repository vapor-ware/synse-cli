package utils

import (
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
	RackID string `json:"rack_id"`
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

func FilterDevices(fn func(Result) bool) chan Result {
	c := make(chan Result)

	go func() {
		for res := range GetDevices() {
			if fn(res) {
				c <- Result{res.Rack, res.Board, res.Device}
			}
		}

		close(c)
	}()

	return c
}

func GetDevices() chan Result {
	c := make(chan Result)

	vc := client.New()
	status := &scanResponse{}
	resp, err := vc.Sling.New().Get(Scanpath).ReceiveSuccess(status)
	if err != nil {
		panic(fmt.Sprintf("status: %s\nerror: %s\n", status, err))
	}
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("status: %s\nerror: %s\n", status, err))
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

	return c
}
