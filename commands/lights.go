package commands

import (
	"fmt"

	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/utils"
)

const lightspath = "led/"
const lightsdevicetype = "led"

// LightsDetails contains the fields returned from a lights query
type LightsDetails struct {
	State      string `json:"led_state"`
	BlinkState string `json:"blink_state"`
	Color      int16  `json:"color"`
}

// LightsResult contains the fields returned and a result object from a lights query
type LightsResult struct {
	utils.Result
	*LightsDetails
}

// ListLights iterates over the complete list of devices and returns blink state,
// color, and state of each `led` device type. Since there may
// be multiple lights per board, each board is also iterated over for each
// device of type `led`.
// Future types may need to be added to this list to accomidate different
// types of led data.
// NOTE: Currently only Chamber LED's support blink state and color. No error
// checking is done on this at the moment.
func ListLights(filter *utils.FilterFunc) ([]LightsResult, error) {
	var devices []utils.Result

	var data []LightsResult

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

	progressBar, pbWriter := utils.ProgressBar(len(devices), "Polling Lights")

	for _, res := range devices {
		lights, _ := GetLights(res)
		data = append(data, lights)
		progressBar.Incr(1)
	}

	utils.ProgressBarStop(pbWriter)
	return data, err
}

// GetLights queries the api for  lights on a specific rack and board. If there
// are no query errors it returns the results.
func GetLights(res utils.Result) (LightsResult, error) {
	lights := &LightsDetails{}
	path := fmt.Sprintf("%s/%s/%s", res.RackID, res.BoardID, res.DeviceID)
	_, err := client.New().Path(lightspath).Get(path).ReceiveSuccess(lights)
	if err != nil {
		return LightsResult{res, lights}, err
	}

	return LightsResult{res, lights}, nil
}

// PrintListLights takes the output from ListLights and pretty prints it into a table.
// Multiple lights are grouped by board, then by rack. Table format is set to not
// auto merge duplicate entries.
func PrintListLights() error {
	filter := &utils.FilterFunc{}
	filter.DeviceType = lightsdevicetype
	filter.FilterFn = func(res utils.Result) bool {
		return res.DeviceType == lightsdevicetype
	}

	header := []string{"Rack", "Board", "Device", "Name", "LED State"}
	lightsList, err := ListLights(filter)
	if err != nil {
		return err
	}

	var data [][]string

	for _, res := range lightsList {
		data = append(data, []string{
			res.RackID,
			res.BoardID,
			res.DeviceID,
			res.DeviceInfo,
			res.State})
	}

	utils.TableOutput(header, data)

	return nil
}

// PrintGetLight takes the output of GetLight and pretty prints it in table form.
// Multiple entries are not merged.
func PrintGetLight(args utils.GetDeviceArgs) error { // FIXME: Change this to "lights"
	filter := &utils.FilterFunc{}
	filter.DeviceType = lightsdevicetype
	filter.RackID = args.RackID
	filter.BoardID = args.BoardID
	filter.FilterFn = func(res utils.Result) bool {
		return res.DeviceType == lightsdevicetype && res.RackID == args.RackID && res.BoardID == args.BoardID
	}

	header := []string{"Rack", "Board", "Device", "Name", "LED State"}
	lightsList, err := ListLights(filter)
	if err != nil {
		return err
	}

	var data [][]string

	for _, res := range lightsList {
		data = append(data, []string{
			res.RackID,
			res.BoardID,
			res.DeviceID,
			res.DeviceInfo,
			res.State})
	}

	utils.TableOutput(header, data)

	return nil
}

// SetLight takes a rack and board id as a locater, as well as a light status.
// The status of the matching light is set to the passed light status.
// Options are: `--state [on/off]`, `--color <color hex>`, `--blink [on/off]`.
func SetLight(rack_id, board_id, light_status string) (string, error) {
	responseData := &LightsDetails{}
	path := fmt.Sprintf("%s/%s/%s/",
		rack_id, board_id, lightsdevicetype)
	resp, err := client.New().Path(lightspath).Path(path).Get(
		light_status).ReceiveSuccess(responseData) // TODO: Add error reporting

	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 { // This is not what I meant by "error reporting"
		return "", err
	}

	return responseData.State, err
}

// PrintSetLight takes input in the form of a rack and board id, command type,
// and command type state. The rack and board id's are used as locators to
// specify a device with type "LED". The light command may be "state", "color",
// or "blink", corresponding to the same action. The command type state is the
// given state to which a specific light command is to be set. For example,
// the light command "blink" may be set to the state "on" or "off". The
// acceptable types differ for each command, and are given in the usage
// documentation for that command.
// Command types and states are specified when running the command by the
// presence of the corresponding flag. For example, the command type "state"
// is given by the flag "--state". The state is given as the argument to this
// flag.
// func PrintSetLight(rack_id, board_id, light_input, light_command string) error {
func PrintSetLight(args utils.SetLightsArgs) error {
	if args.State != "" {
		status, err := SetLight(args.RackID, args.BoardID, args.State)
		fmt.Println(status)
		return err
	}

	if args.Color != "" {
		light_action := fmt.Sprintf("state/color/%s", args.Color) // Might need this to be a nonstring input
		status, err := SetLight(args.RackID, args.BoardID, light_action)
		fmt.Println(status)
		return err
	}

	if args.Blink != "" {
		light_action := fmt.Sprintf("state/blink_state/%s", args.Blink)
		status, err := SetLight(args.RackID, args.BoardID, light_action)
		fmt.Println(status)
		return err
	}

	return nil
}
