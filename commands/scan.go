package commands

import (
	"net/http"

	"github.com/vapor-ware/vesh/client"
	"github.com/vapor-ware/vesh/utils"
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

// TODO: walkRacks is not yet implemented.
func walkRacks(sr *scanResponse) {

}

// TODO: walkBoards is not yet implemented.
func walkBoards(sr *scanResponse) {

}

// Scan polls the infrastructure (using the `/scan` endpoint) and assigns the
// responses to the appropriate fields in the scanResponse struct. Because the
// json response contains multiple nested levels of data, each level is walked
// to populate "bottom level" data.
//
// Because walking the full tree can take some time, a progress bar is displayed
// during the scan process.
// NOTE: Printing output is part of this function. To access scan results
// internally, utils.UtilScanOnly should be used.
func Scan(vc *client.VeshClient) error {
	var data [][]string

	filter := func(res utils.Result) bool {
		return true
	}

	for res := range utils.FilterDevices(filter) {
		data = append(data, []string{
			res.RackID,
			res.BoardID,
			res.DeviceID,
			res.DeviceInfo,
			res.DeviceType})
	}

	header := []string{"Rack", "Board", "Device ID", "Device Info", "Device Type"}
	utils.TableOutput(header, data)

	return nil
}

func writetable() {

}
