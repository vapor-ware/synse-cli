package commands

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/utils"
)

// fanpath specifies the endpoint to query for fan status
const fanpath = "fan/"

// fandevicetype specifies the device type to query for fan status
const fandevicetype = "fan"

// FanDetails contains the response data for a fan object
type FanDetails struct {
	Health   string   `json:"health"`
	SpeedRPM float64  `json:"speed_rpm"`
	States   []string `json:"states"`
}

// FanResult combines the standard utils.Result data with FanDetails
type FanResult struct {
	utils.Result
	*FanDetails
}

// ListFan iterates over the complete list of devices and returns health,
// speed (rpm), and state of each `fan` device type. Since there may
// be multiple fans per board, each board is also iterated over for each
// device of type `fan`.
// Future types may need to be added to this list to accomidate different
// types of fan data.
func ListFan(filter *utils.FilterFunc) ([]FanResult, error) {
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
		fan, err := GetFan(res)
		if err != nil {
			return data, err
		}
		data = append(data, FanResult{res, fan})
		progressBar.Incr(1)
	}

	utils.ProgressBarStop(pbWriter)
	return data, err
}

// GetFan queries the api for any fan located on a specific rack and board. If
// there is no query error, it returns the fan details associated with that board.
func GetFan(res utils.Result) (*FanDetails, error) {
	fan := &FanDetails{}
	failure := new(client.ErrorResponse)
	path := fmt.Sprintf("%s/%s/%s", res.RackID, res.BoardID, res.DeviceID)
	resp, err := client.New().Path(fanpath).Get(path).Receive(fan, failure)
	if err != nil {
		return fan, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%+v\n%+v", failure.HTTPCode, failure.Message)
	}

	return fan, nil
}

// PrintListFan takes the output from ListFan and pretty prints it into a table.
// Multiple fans are grouped by board, then by rack. Table format is set to not
// auto merge duplicate entries.
func PrintListFan() error {
	filter := &utils.FilterFunc{}
	filter.FilterFn = func(res utils.Result) bool {
		return res.DeviceType == fandevicetype
	}

	header := []string{"Rack", "Board", "Device", "Name", "Fan Speed (RPM)"}
	fanList, err := ListFan(filter)
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
func PrintGetFan(args utils.GetDeviceArgs) error {
	filter := &utils.FilterFunc{}
	filter.DeviceType = fandevicetype
	filter.RackID = args.RackID
	filter.BoardID = args.BoardID
	filter.FilterFn = func(res utils.Result) bool {
		return res.DeviceType == fandevicetype && res.RackID == args.RackID && res.BoardID == args.BoardID
	}

	header := []string{"Rack", "Board", "Device", "Name", "Health", "Speed (RPM)", "States"}
	fanList, err := ListFan(filter) // Add error reporting
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
