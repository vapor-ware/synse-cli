package commands

import (
	"fmt"
	"strings"
	"net/http"
	"errors"

	"github.com/vapor-ware/vesh/client"
	"github.com/vapor-ware/vesh/utils"
)

const fanpath = "fan/"
const fandevicetype = "fan_speed"

type FanDetails struct {
	Health   string   `json:"health"`
	SpeedRPM float64  `json:"speed_rpm"`
	States   []string `json:"states"`
}

type FanResult struct {
	utils.Result
	*FanDetails
}

// ListFan iterates over the complete list of devices and returns health,
// speed (rpm), and state of each `fan_speed` device type. Since there may
// be multiple fans per board, each board is also iterated over for each
// device of type `fan_speed`.
// Future types may need to be added to this list to accomidate different
// types of fan data.
func ListFan(vc *client.VeshClient, filter *utils.FilterFunc) ([]FanResult, error) {
	var devices []utils.Result

	var data []FanResult

	fil, err := utils.FilterDevices(filter)
	if err != nil {
		return data, err
	}
	for res := range fil {
		if res.Error != nil {
			return data, res.Error
		}
		devices = append(devices, res.Result)
	}

	progressBar, pbWriter := utils.ProgressBar(len(devices), "Polling Fans")

	for _, res := range devices {
		fan, err := GetFan(vc, res)
		if err != nil {
			return data, err
		}
		data = append(data, FanResult{res, fan})
		progressBar.Incr(1)
	}

	utils.ProgressBarStop(pbWriter)
	return data, err
}

func GetFan(vc *client.VeshClient, res utils.Result) (*FanDetails, error) {
	fan := &FanDetails{}
	failure := new(client.ErrorResponse)
	path := fmt.Sprintf("%s/%s/%s", res.RackID, res.BoardID, res.DeviceID)
	resp, err := vc.Sling.New().Path(fanpath).Get(path).Receive(fan, failure)
	if err != nil {
		return fan, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("%+v\n%+v", failure.HttpCode, failure.Message))
	}

	return fan, nil
}

// PrintListFan takes the output from ListFan and pretty prints it into a table.
// Multiple fans are grouped by board, then by rack. Table format is set to not
// auto merge duplicate entries.
func PrintListFan(vc *client.VeshClient) error {
	filter := &utils.FilterFunc{}
	filter.FilterFn = func(res utils.Result) bool {
		return res.DeviceType == fandevicetype
	}

	header := []string{"Rack", "Board", "Device", "Name", "Fan Speed (RPM)"}
	fanList, err := ListFan(vc, filter)
	if err != nil {
		return err
	}

	var data [][]string

	for _, res := range fanList {
		data = append(data, []string{
			res.RackID,
			res.BoardID,
			res.DeviceID,
			res.DeviceInfo,
			fmt.Sprintf("%.0f", res.SpeedRPM)})
	}

	utils.TableOutput(header, data)

	return err
}

// PrintGetFan takes the output of GetFan and pretty prints it in table form.
// Multiple entries are not merged.
func PrintGetFan(vc *client.VeshClient, rack_id, board_id string) error {
	filter := &utils.FilterFunc{}
	filter.Path = fanpath
	filter.DeviceType = fandevicetype
	filter.RackID = rack_id
	filter.BoardID = board_id
	filter.FilterFn = func(res utils.Result) bool {
		return res.DeviceType == fandevicetype && res.RackID == rack_id && res.BoardID == board_id
	}

	header := []string{"Rack", "Board", "Device", "Name", "Health", "Speed (RPM)", "States"}
	fanList, err := ListFan(vc, filter) // Add error reporting
	if err != nil {
		return err
	}

	var data [][]string

	for _, res := range fanList {
		data = append(data, []string{
			res.RackID,
			res.BoardID,
			res.DeviceID,
			res.DeviceInfo,
			res.Health,
			fmt.Sprintf("%.0f", res.SpeedRPM),
			strings.Join(res.States, ",")})
	}

	utils.TableOutput(header, data)

	return nil
}
