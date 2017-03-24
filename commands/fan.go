package commands

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/vapor-ware/vesh/client"
	"github.com/vapor-ware/vesh/utils"

	"github.com/olekukonko/tablewriter"
	"github.com/gosuri/uiprogress"
)

const fanpath = "fan/"
const fandevicetype = "fan_speed"

type fanResponse struct {
	Health   string   `json:"health"`
	SpeedRPM float64  `json:"speed_rpm"`
	States   []string `json:"states"`
}

// ListFan iterates over the complete list of devices and returns health,
// speed (rpm), and state of each `fan_speed` device type. Since there may
// be multiple fans per board, each board is also iterated over for each
// device of type `fan_speed`.
// Future types may need to be added to this list to accomidate different
// types of fan data.
func ListFan(vc *client.VeshClient) ([][]string, error) {
	uiprogress.Start()
	progressBar:= uiprogress.AddBar(utils.TotalElemsNum())
	progressBar.AppendCompleted()
	progressBar.PrependElapsed()
	scanResponse, _ := utils.UtilScanOnly() // Add error reporting
	scanResponsePtr := reflect.ValueOf(&scanResponse.Racks)
	scanResponseValuePtr := scanResponsePtr.Elem()
	fulltable := make([][]string, 0)
	totalruns := 0
	for i := 0; i < scanResponseValuePtr.Len(); i++ {
		boardsPtr := reflect.ValueOf(&scanResponse.Racks[i].Boards)
		boardsValuePtr := boardsPtr.Elem()
		for j := 0; j < boardsValuePtr.Len(); j++ {
			devicePtr := reflect.ValueOf(&scanResponse.Racks[i].Boards[j].Devices)
			devicesValuePtr := devicePtr.Elem()
			for k := 0; k < devicesValuePtr.Len(); k++ {
				deviceTypePtr := reflect.ValueOf(&scanResponse.Racks[i].Boards[j].Devices[k].DeviceType)
				deviceTypeValuePtr := deviceTypePtr.Elem()
				progressBar.Incr()
				if deviceTypeValuePtr.String() == fandevicetype { // This may need to be expanded to other types
					tablerow := make([]string, 0)
					rack_id := scanResponse.Racks[i].RackID
					board_id := scanResponse.Racks[i].Boards[j].BoardID
					device_id := scanResponse.Racks[i].Boards[j].Devices[k].DeviceID
					tablerow = append(tablerow, rack_id)
					tablerow = append(tablerow, board_id)
					tablerow = append(tablerow, device_id)
					responseData := &fanResponse{}
					resp, err := vc.Sling.New().Path(fanpath).Path(rack_id + "/").Path(board_id + "/").Get(device_id).ReceiveSuccess(responseData) // Add error reporting
					if resp.StatusCode != 200 {                                                                                                    // This is not what I meant by "error reporting"
						fmt.Println(vc)
						fmt.Println(resp)
						return nil, err
					}
					tablerow = append(tablerow, strconv.FormatFloat(responseData.SpeedRPM, 'G', -1, 64))
					fulltable = append(fulltable, nil)
					fulltable[totalruns] = make([]string, 0)
					fulltable[totalruns] = append(fulltable[totalruns], tablerow...)
					totalruns++
				}
			}
		}
	}
	uiprogress.Stop()
	return fulltable, nil
}

// PrintListFan takes the output from ListFan and pretty prints it into a table.
// Multiple fans are grouped by board, then by rack. Table format is set to not
// auto merge duplicate entries.
func PrintListFan(vc *client.VeshClient) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rack", "Board", "Device", "Fan Speed"})
	table.SetBorder(false)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAutoMergeCells(false)
	fmt.Println("Polling fans. This may take some time...")
	fanList, _ := ListFan(vc) // Add error reporting
	table.AppendBulk(fanList)
	table.Render()
	return nil
}

// GetFan takes a rack and board id as a locator and returns the health,
// speed (rpm), and state of all fans on that board.
func GetFan(vc *client.VeshClient, rack_id, board_id string) ([][]string, error) {
	scanResponse, scanerr := utils.UtilScanOnly() // Add error reporting
	rackidint, _ := strconv.Atoi(rack_id)
	boardidint, _ := strconv.Atoi(board_id)
	devicePtr := reflect.ValueOf(&scanResponse.Racks[utils.RackIDtoElem(rackidint)].Boards[utils.BoardIDtoElem(boardidint)].Devices)
	deviceValuePtr := devicePtr.Elem()
	fulltable := make([][]string, 0)
	totalruns := 0
	for i := 0; i < deviceValuePtr.Len(); i++ {
		if scanResponse.Racks[utils.RackIDtoElem(rackidint)].Boards[utils.BoardIDtoElem(boardidint)].Devices[i].DeviceType == fandevicetype {
			device_id := scanResponse.Racks[utils.RackIDtoElem(rackidint)].Boards[utils.BoardIDtoElem(boardidint)].Devices[i].DeviceID
			responseData := &fanResponse{}
			resp, err := vc.Sling.New().Path(fanpath).Path(rack_id + "/").Path(board_id + "/").Get(device_id).ReceiveSuccess(responseData) // Add error reporting
			if resp.StatusCode != 200 {                                                                                                    // This is not what I meant by "error reporting"
				return nil, err
			}
			tablerow := make([]string, 0)
			tablerow = append(tablerow, rack_id, board_id, device_id, responseData.Health, strconv.FormatFloat(responseData.SpeedRPM, 'G', -1, 64))
			tablerow = append(tablerow, responseData.States...)
			fulltable = append(fulltable, nil)
			fulltable[totalruns] = make([]string, 0)
			fulltable[totalruns] = append(fulltable[totalruns], tablerow...)
			totalruns++
		}
	}
	return fulltable, scanerr
}

// PrintGetFan takes the output of GetFan and pretty prints it in table form.
// Multiple entries are not merged.
func PrintGetFan(vc *client.VeshClient, rack_id, board_id string) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rack", "Board", "Device", "Health", "Speed (RPM)", "States"})
	table.SetBorder(false)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAutoMergeCells(false)
	fanStatus, _ := GetFan(vc, rack_id, board_id) // Add error reporting
	table.AppendBulk(fanStatus)
	table.Render()
	return nil
}
