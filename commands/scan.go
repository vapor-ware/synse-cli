package commands

import (
	"fmt"
	"net/http"
	"os"
	"reflect"

	"github.com/vapor-ware/vesh/client"
	"github.com/vapor-ware/vesh/utils"

	"github.com/olekukonko/tablewriter"
	"github.com/gosuri/uiprogress"
)

const Scanpath = "scan"

// scanResponse struct holds the response values from a `/scan` operation.
// While it does not contain the complete set of information available, it does
// contain a complete list of the available assets, including racks and boards.
// The structure mirrors the json struture of response from `/scan` and values
// are assigned to appropriate sub structs.
type scanResponse struct {
	Racks []struct {
		Boards []struct {
			BoardID     string   `json:"board_id"`
			Hostnames   []string `json:"hostnames"`
			IPAddresses []string `json:"ip_addresses"`
			Devices     []struct {
				DeviceID   string `json:"device_id"`
				DeviceInfo string `json:"device_info"`
				DeviceType string `json:"device_type"`
			} `json:"devices"`
		} `json:"boards"`
		RackID string `json:"rack_id"`
	} `json:"racks"`
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
	totaltouched := 0
	uiprogress.Start()
	progressBar := uiprogress.AddBar(utils.TotalElemsNum())
	data := make([][]string, 10000)
	racksPtr := reflect.ValueOf(&status.Racks)
	racksValuePtr := racksPtr.Elem()
	for i := 0; i < racksValuePtr.Len(); i++ {
		boardsPtr := reflect.ValueOf(&status.Racks[i].Boards)
		boardsValuePtr := boardsPtr.Elem()
		for j := 0; j < boardsValuePtr.Len(); j++ {
			devicesPtr := reflect.ValueOf(&status.Racks[i].Boards[j].Devices)
			devicesValuePtr := devicesPtr.Elem()
			for k := 0; k < devicesValuePtr.Len(); k++ {
				devicePtr := reflect.ValueOf(&status.Racks[i].Boards[j].Devices[i])
				deviceValuePtr := devicePtr.Elem()
				tablerow := make([]string, 0)
				data = append(data, nil)
				rack_id := status.Racks[i].RackID
				board_id := status.Racks[i].Boards[j].BoardID
				tablerow = append(tablerow, rack_id, board_id)
				for l := 0; l < deviceValuePtr.NumField(); l++ {
					tablerow = append(tablerow, deviceValuePtr.Field(l).String())
					progressBar.Incr()
				}
				data[totaltouched] = append(data[totaltouched], tablerow...)
				totaltouched++
			}
		}
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rack", "Board", "Device ID", "Device Info", "Device Type"})
	table.SetBorder(false)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
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
