package commands

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/vapor-ware/vesh/client"
	"github.com/vapor-ware/vesh/utils"

	"github.com/olekukonko/tablewriter"
	"github.com/sethgrid/multibar"
)

const powerpath = "power/"
const device_id = "power"

type powerResponse struct {
	InputPower  float64 `json:"input_power"`
	OverCurrent bool    `json:"over_current"`
	PowerOK     bool    `json:"power_ok"`
	PowerStatus string  `json:"power_status"`
}

func ListPower(vc *client.VeshClient) ([][]string, error) {
	scanResponse, _ := ScanOnly(vc) // Add error reporting
	scanResponsePtr := reflect.ValueOf(&scanResponse.Racks)
	scanResponseValuePtr := scanResponsePtr.Elem()
	fulltable := make([][]string, 0)
	totalruns := 0
	totaltouched := 0
	progressBar, _ := multibar.New()
	go progressBar.Listen()
	polling := progressBar.MakeBar(utils.TotalElemsNum(), "Polling power states") // Should use total number of boards since we're assuming only one 'power' device per board
	for i := 0; i < scanResponseValuePtr.Len(); i++ {
		boardsPtr := reflect.ValueOf(&scanResponse.Racks[i].Boards)
		boardsValuePtr := boardsPtr.Elem()
		totaltouched++
		for j := 0; j < boardsValuePtr.Len(); j++ {
			tablerow := make([]string, 0)
			rack_id := scanResponse.Racks[i].RackID
			board_id := scanResponse.Racks[i].Boards[j].BoardID
			tablerow = append(tablerow, rack_id)
			tablerow = append(tablerow, scanResponse.Racks[i].Boards[j].BoardID) // Switch this to the variable
			polling(totaltouched)
			responseData := &powerResponse{}
			resp, err := vc.Sling.New().Path(powerpath).Path(rack_id + "/").Path(board_id + "/").Get(device_id).ReceiveSuccess(responseData) // Add error reporting
			if resp.StatusCode != 200 {                                                                                                      // This is not what I meant by "error reporting"
				return nil, err
			}
			tablerow = append(tablerow, strconv.FormatFloat(responseData.InputPower, 'G', -1, 64))
			tablerow = append(tablerow, strconv.FormatBool(responseData.PowerOK))
			fulltable = append(fulltable, nil)
			fulltable[totalruns] = make([]string, 0)
			fulltable[totalruns] = append(fulltable[totalruns], tablerow...)
			totalruns++
		}
	}
	return fulltable, nil
	//return nil, scanerr //fix with proper error
}

func PrintListPower(vc *client.VeshClient) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rack", "Board", "Input Power", "Power Ok?"})
	table.SetBorder(false)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAutoMergeCells(true)
	powerList, _ := ListPower(vc) // Add error reporting
	table.AppendBulk(powerList)
	table.Render()
	return nil
}

func GetPower(vc *client.VeshClient, rack_id, board_id string) ([]string, error) {
	responseData := &powerResponse{}
	resp, err := vc.Sling.New().Path(powerpath).Path(rack_id + "/").Path(board_id + "/").Get(device_id).ReceiveSuccess(responseData) // Add error reporting
	if resp.StatusCode != 200 {                                                                                                      // This is not what I meant by "error reporting"
		return nil, err
	}
	tablerow := make([]string, 0)
	tablerow = append(tablerow, rack_id, board_id, strconv.FormatFloat(responseData.InputPower, 'G', -1, 64), strconv.FormatBool(responseData.OverCurrent), strconv.FormatBool(responseData.PowerOK), responseData.PowerStatus)
	return tablerow, nil
}

func PrintGetPower(vc *client.VeshClient, rack_id, board_id string) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Rack", "Board", "Input Power", "Over Current?", "Power Ok?", "Power Status"})
	table.SetBorder(false)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	tablerow, err := GetPower(vc, rack_id, board_id)
	table.Append(tablerow)
	table.Render()
	return err // Add error reporting
}

func SetPower(vc *client.VeshClient, rack_id, board_id, power_status string) (string, error) {
	responseData := &powerResponse{}
	resp, err := vc.Sling.New().Path(powerpath).Path(rack_id + "/").Path(board_id + "/").Path(device_id + "/").Get(power_status).ReceiveSuccess(responseData) // Add error reporting
	if resp.StatusCode != 200 {                                                                                                                               // This is not what I meant by "error reporting"
		return "", err
	}
	if err == nil && power_status == "cycle" { // This should check if successful
		return power_status, err
	}
	return responseData.PowerStatus, err
}

func PrintSetPower(vc *client.VeshClient, rack_id, board_id, power_status string) error {
	status, err := SetPower(vc, rack_id, board_id, power_status)
	if err == nil && status == "cycle" {
		fmt.Printf("Power successfully %sd\n", status)
	}
	if err == nil && status != "cycle" {
		fmt.Printf("Power set to %s\n", status)
	}
	return err
}
