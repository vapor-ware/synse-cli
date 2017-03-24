package commands

import (
	"fmt"
	"strings"

	"github.com/vapor-ware/vesh/client"
	"github.com/vapor-ware/vesh/utils"

	"github.com/gosuri/uiprogress"
)

const fanpath = "fan/"
const fandevicetype = "fan_speed"

type FanDetails struct {
	Health   string   `json:"health"`
	SpeedRPM float64  `json:"speed_rpm"`
	States   []string `json:"states"`
}

type FanResult struct {
	res         utils.Result
	fan         *FanDetails
}

// ListFan iterates over the complete list of devices and returns health,
// speed (rpm), and state of each `fan_speed` device type. Since there may
// be multiple fans per board, each board is also iterated over for each
// device of type `fan_speed`.
// Future types may need to be added to this list to accomidate different
// types of fan data.
func ListFan(vc *client.VeshClient, filter func(res utils.Result) bool) ([]FanResult, error) {
	var devices []utils.Result

	var data []FanResult

	for res := range utils.FilterDevices(filter) {
		devices = append(devices, res)
	}

	uiprogress.Start()
	progressBar := uiprogress.AddBar(len(devices))
	progressBar.AppendCompleted()
	progressBar.PrependElapsed()

	for _, res := range devices {
		fan, _ := GetFan(vc, res)
		data = append(data, FanResult{ res, fan })
		progressBar.Incr()
	}

	return data, nil
}

func GetFan(vc *client.VeshClient, res utils.Result) (*FanDetails, error) {
	fan := &FanDetails{}
	path := fmt.Sprintf("%s/%s/%s", res.Rack.RackID, res.Board.BoardID, res.Device.DeviceID)
	_, err := vc.Sling.New().Path(fanpath).Get(path).ReceiveSuccess(fan)
	if err != nil {
		return fan, err
	}

	return fan, nil
}

// PrintListFan takes the output from ListFan and pretty prints it into a table.
// Multiple fans are grouped by board, then by rack. Table format is set to not
// auto merge duplicate entries.
func PrintListFan(vc *client.VeshClient) error {
	filter := func(res utils.Result) bool {
		return res.Device.DeviceType == fandevicetype
	}

	header := []string{"Rack", "Board", "Device", "Name", "Fan Speed"}
	fanList, _ := ListFan(vc, filter)

	var data [][]string

	for _, res := range fanList {
			data = append(data, []string{
				res.res.Rack.RackID,
				res.res.Board.BoardID,
				res.res.Device.DeviceID,
				res.res.Device.DeviceInfo,
				fmt.Sprintf("%.0f", res.fan.SpeedRPM)})
	}

	utils.TableOutput(header, data)

	return nil
}

// PrintGetFan takes the output of GetFan and pretty prints it in table form.
// Multiple entries are not merged.
func PrintGetFan(vc *client.VeshClient, rack_id, board_id string) error {
	filter := func(res utils.Result) bool {
		return res.Device.DeviceType == fandevicetype && res.Rack.RackID == rack_id && res.Board.BoardID == board_id
	}

	header := []string{"Rack", "Board", "Device", "Name", "Health", "Speed (RPM)", "States"}
	fanList, _ := ListFan(vc, filter) // Add error reporting

	var data [][]string

	for _, res := range fanList {
		data = append(data, []string{
			res.res.Rack.RackID,
			res.res.Board.BoardID,
			res.res.Device.DeviceID,
			res.res.Device.DeviceInfo,
			res.fan.Health,
			fmt.Sprintf("%.0f", res.fan.SpeedRPM),
			strings.Join(res.fan.States, ",")})
	}

	utils.TableOutput(header, data)

	return nil
}
