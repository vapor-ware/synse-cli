package commands

import (
	"github.com/vapor-ware/synse-cli/utils"
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
	RackID string  `json:"rack_id"`
}

// Board contains the top level objects for a board
type Board struct {
	BoardID     string   `json:"board_id"`
	Hostnames   []string `json:"hostnames"`
	IPAddresses []string `json:"ip_addresses"`
	Devices     []Device `json:"devices"`
}

// Device contains the response values for a specific device
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
func Scan() error {
	var data [][]string

	filter := &utils.FilterFunc{}
	filter.FilterFn = func(res utils.Result) bool {
		return true
	}

	fil, err := utils.FilterDevices(filter)
	if err != nil {
		return err
	}
	for res := range fil {
		if res.Error != nil {
			return res.Error
		}
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
