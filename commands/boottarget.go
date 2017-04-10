package commands

import (
	"fmt"
	"net/http"

	"github.com/vapor-ware/vesh/client"
	"github.com/vapor-ware/vesh/utils"
)

const bootpath = "boot_target/"
const bootdevicetype = "system"

type BootTargetDetails struct {
	Target string `json:"target"`
	Status string `json:"status"`
}

type BootTargetResult struct {
	utils.Result
	*BootTargetDetails
}

func ListBootTarget(filter *utils.FilterFunc) ([]BootTargetResult, error) {
	var devices []utils.Result

	var data []BootTargetResult

	fil, err := utils.FilterDevices(filter)
	if err != nil {
		return data, nil
	}
	for res := range fil {
		if res.Error != nil {
			return data, res.Error
		}
		devices = append(devices, res.Result)
	}

	progressBar, pbWriter := utils.ProgressBar(len(devices), "Polling Boot Targets")

	for _, res := range devices {
		boottarget, _ := GetBootTarget(res)
		data = append(data, BootTargetResult{res, boottarget})
		progressBar.Incr(1)
	}

	utils.ProgressBarStop(pbWriter)
	return data, err
}

func PrintListBootTarget() error {
	filter := &utils.FilterFunc{}
	filter.DeviceType = bootdevicetype
	filter.FilterFn = func(res utils.Result) bool {
		return res.DeviceType == bootdevicetype
	}

	header := []string{"Rack", "Board", "Device", "Device Type", "Boot Target"}
	targetList, err := ListBootTarget(filter)
	if err != nil {
		return err
	}

	var data [][]string

	for _, res := range targetList {
		data = append(data, []string{
			res.RackID,
			res.BoardID,
			res.DeviceID,
			res.DeviceInfo,
			res.Target})
	}

	utils.TableOutput(header, data)

	return nil
}

// GetCurrentBootTarget takes a rack and board id as locators and returns the
// current set boot target for the `system` device.
func GetBootTarget(res utils.Result) (*BootTargetDetails, error) {
	boottarget := &BootTargetDetails{}
	path := fmt.Sprintf("%s/%s/%s", res.RackID, res.BoardID, res.DeviceID)
	_, err := client.New().Path(bootpath).Get(path).ReceiveSuccess(boottarget)
	if err != nil {
		return boottarget, err
	}

	return boottarget, nil
}

// PrintGetCurrentBootTarget calls GetCurrentBootTarget and pretty prints the
// result.
func PrintGetBootTarget(args utils.GetDeviceArgs) error {
	filter := &utils.FilterFunc{}
	filter.DeviceType = bootdevicetype
	filter.RackID = args.RackID
	filter.BoardID = args.BoardID
	filter.FilterFn = func(res utils.Result) bool {
		return res.DeviceType == bootdevicetype && res.RackID == args.RackID && res.BoardID == args.BoardID
	}

	header := []string{"Rack", "Board", "Device", "Device Type", "Boot Target", "Status"}
	targetList, err := ListBootTarget(filter)
	if err != nil {
		return err
	}

	var data [][]string

	for _, res := range targetList {
		data = append(data, []string{
			res.RackID,
			res.BoardID,
			res.DeviceID,
			res.DeviceInfo,
			res.Target,
			res.Status})
	}

	utils.TableOutput(header, data)

	return nil
}

// SetCurrentBootTarget takes a rack and board id as locators as well as a
// boot target. The matching `system` device's boot target is set to the passed
// boot target.
// Options are: `no-override`, `hdd`, `pxe`.
func SetBootTarget(args utils.SetBootTargetArgs) error {
	status := &BootTargetDetails{}
	path := fmt.Sprintf("%s/%s/%s/%s/", bootpath, args.RackID, args.BoardID, bootdevicetype)
	resp, err := client.New().Path(path).Get(args.Value).ReceiveSuccess(status)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return err
	}
	fmt.Println(status.Target, status.Status)
	return nil
}
