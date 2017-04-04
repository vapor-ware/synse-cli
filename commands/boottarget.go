package commands

import (
	"fmt"
	"net/http"

	"github.com/vapor-ware/vesh/client"
	"github.com/vapor-ware/vesh/utils"
)

const bootpath = "boot_target/"

var bootdevicetype = "system"

type boottargetresponse struct {
	Target string `json:"target"`
	status string `json:"status"`
}

// GetCurrentBootTarget takes a rack and board id as locators and returns the
// current set boot target for the `system` device.
func GetCurrentBootTarget(rack_id, board_id string) (string, error) {
	status := &boottargetresponse{}
	failure := new(client.ErrorResponse)

	resp, err := client.New().Path(
		fmt.Sprintf("%s/%s/%s/", bootpath, rack_id, board_id)).Get(
		bootdevicetype).Receive(status, failure)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return failure.Message, err
	}
	return status.Target, nil
}

// PrintGetCurrentBootTarget calls GetCurrentBootTarget and pretty prints the
// result.
func PrintGetCurrentBootTarget(args utils.GetDeviceArgs) error {
	bootTarget, err := GetCurrentBootTarget(args.RackID, args.BoardID)
	if err != nil {
		return err
	}
	fmt.Println(bootTarget)
	return nil
}

// SetCurrentBootTarget takes a rack and board id as locators as well as a
// boot target. The matching `system` device's boot target is set to the passed
// boot target.
// Options are: `no-override`, `hdd`, `pxe`.
func SetCurrentBootTarget(args utils.SetBootTargetArgs) error {
	status := &boottargetresponse{}
	path := fmt.Sprintf("%s/%s/%s/%s/", bootpath, args.RackID, args.BoardID, bootdevicetype)
	resp, err := client.New().Path(path).Get(args.Value).ReceiveSuccess(status)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return err
	}
	fmt.Println(status.Target, status.status)
	return nil
}
