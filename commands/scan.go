package commands

import (
	"fmt"
	"net/http"
	"os"

	"github.com/vapor-ware/vesh/client"

	"github.com/fatih/structs"
	"github.com/olekukonko/tablewriter"
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
func Scan(vc *client.VeshClient) (*scanResponse, error) {
	status := &scanResponse{}
	resp, err := vc.Sling.New().Get(Scanpath).ReceiveSuccess(status)
	if err != nil {
		return status, err
	}
	if resp.StatusCode != http.StatusOK {
		return status, err
	}
	fmt.Println("API reported status ok")

	var data [][]string

	for _, rack := range structs.New(status).Field("Racks").Value().([]Rack) {
		for _, board := range structs.New(rack).Field("Boards").Value().([]Board) {
			for _, device := range structs.New(board).Field("Devices").Value().([]Device) {
				data = append(data, []string{
					rack.RackID,
					board.BoardID,
					device.DeviceID,
					device.DeviceInfo,
					device.DeviceType})
			}
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Rack", "Board", "Device ID", "Device Info", "Device Type"})
	table.SetBorder(false)
	table.SetBorders(tablewriter.Border{
		Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()

	return status, nil
}

// ScanOnly returns the results of a scan without any formatting or printing.
// NOTE: This function may be removed in favor of util.UtilScanOnly.
func ScanOnly(vc *client.VeshClient) (*scanResponse, error) {
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

func writetable() {

}
