package commands

import (
	"fmt"
	"net/http"

	"github.com/vapor-ware/vesh/client"
	//"github.com/olekukonko/tablewriter"
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
func PrintGetCurrentBootTarget(rack_id, board_id string) error {
	bootTarget, err := GetCurrentBootTarget(rack_id, board_id)
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
func SetCurrentBootTarget(rack_id, board_id, boot_target string) error {
	status := &boottargetresponse{}
	path := fmt.Sprintf("%s/%s/%s/%s/", bootpath, rack_id, board_id, bootdevicetype)
	resp, err := client.New().Path(path).Get(boot_target).ReceiveSuccess(status)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return err
	}
	fmt.Println(status.Target, status.status)
	return nil
}
