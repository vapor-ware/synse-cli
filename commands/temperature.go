package commands

import (
	"fmt"
	"strings"

	"github.com/vapor-ware/vesh/client"
	"github.com/vapor-ware/vesh/utils"
)

const temperaturepath = "temperature/"
const temperaturedevicetype = "temperature"

type TempDetails struct {
	Health       string   `json:"health"`
	States       []string `json:"states"`
	TemperatureC float64  `json:"temperature_c"`
}

type TempResult struct {
	utils.Result
	*TempDetails
}

// ListTemp iterates over the complete list of devices and returns health,
// states, and temperature (celcius) for each `temperature` device type.
func ListTemp(vc *client.VeshClient, filter func(res utils.Result) bool) ([]TempResult, error) {
	var devices []utils.Result

	var data []TempResult

	fil, err := utils.FilterDevices(filter)
	for res := range fil {
		devices = append(devices, res)
	}

	progressBar, pbWriter := utils.ProgressBar(len(devices), "Polling Temperatures")

	for _, res := range devices {
		temp, _ := GetTemp(vc, res)
		data = append(data, temp)
		progressBar.Incr(1)
	}

	utils.ProgressBarStop(pbWriter)
	return data, err
}

func GetTemp(vc *client.VeshClient, res utils.Result) (TempResult, error) {
	temp := &TempDetails{}
	path := fmt.Sprintf("%s/%s/%s", res.RackID, res.BoardID, res.DeviceID)
	_, err := vc.Sling.New().Path("read/").Path(temperaturepath).Get(path).ReceiveSuccess(temp)
	if err != nil {
		return TempResult{res, temp}, err
	}

	return TempResult{res, temp}, nil
}

// PrintListTemp takes the output from ListTemp and pretty prints it into a table.
// Multiple temperature readings are grouped by board, then by rack. Table format
// is set to not auto merge duplicate entries.
func PrintListTemp(vc *client.VeshClient) error {
	filter := func(res utils.Result) bool {
		return res.DeviceType == temperaturedevicetype
	}

	header := []string{"Rack", "Board", "Name", "Temperature in C"}
	tempList, _ := ListTemp(vc, filter)

	var data [][]string

	for _, res := range tempList {
		data = append(data, []string{
			res.RackID,
			res.BoardID,
			res.DeviceInfo,
			fmt.Sprintf("%.2f", res.TemperatureC)})
	}

	utils.TableOutput(header, data)

	return nil
}

// PrintGetTemp takes the output of GetTemp and pretty prints it in table form.
// Multiple entries are not merged.
func PrintGetTemp(vc *client.VeshClient, rack_id, board_id string) error {
	filter := func(res utils.Result) bool {
		return res.DeviceType == temperaturedevicetype && res.RackID == rack_id && res.BoardID == board_id
	}

	header := []string{"Rack", "Board", "Device", "Name", "Health", "Temperature in C", "States"}
	tempList, _ := ListTemp(vc, filter) // Add error reporting

	var data [][]string

	for _, res := range tempList {
		data = append(data, []string{
			res.RackID,
			res.BoardID,
			res.DeviceID,
			res.DeviceInfo,
			res.Health,
			fmt.Sprintf("%.2f", res.TemperatureC),
			strings.Join(res.States, ",")})
	}

	utils.TableOutput(header, data)

	return nil
}
