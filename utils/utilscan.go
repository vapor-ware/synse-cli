// Utils package provides useful functions for internal use.
package utils

import (
	"net/http"

	"github.com/vapor-ware/vesh/client"
)

// I DON'T LIKE THIS AT ALL

// const Scanpath = "scan"

// scanResponse struct holds the response values from a `/scan` operation.
// While it does not contain the complete set of information available, it does
// contain a complete list of the available assets, including racks and boards.
// The structure mirrors the json struture of response from `/scan` and values
// are assigned to appropriate sub structs.
// type scanResponse struct {
// 	Racks []struct {
// 		Boards []struct {
// 			BoardID     string   `json:"board_id"`
// 			Hostnames   []string `json:"hostnames"`
// 			IPAddresses []string `json:"ip_addresses"`
// 			Devices     []struct {
// 				DeviceID   string `json:"device_id"`
// 				DeviceInfo string `json:"device_info"`
// 				DeviceType string `json:"device_type"`
// 			} `json:"devices"`
// 		} `json:"boards"`
// 		RackID string `json:"rack_id"`
// 	} `json:"racks"`
// }

// UtilScanOnly polls the infrastructure (using the `/scan` endpoint) and assigns the
// responses to the appropriate fields in the scanResponse struct. Because the
// json response contains multiple nested levels of data, each level is walked
// to populate "bottom level" data.
//
// UtilScanOnly is intended to the equivalent of commands.Scan for internal
// use. No formatting or printing is done on output data.
func UtilScanOnly() (*scanResponse, error) {
	vc := client.New()
	status := &scanResponse{}
	resp, err := vc.Sling.New().Get(Scanpath).ReceiveSuccess(status)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return status, err
	}
	return status, nil
}
