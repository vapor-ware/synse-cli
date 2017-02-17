package commands

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/vapor-ware/vesh/client"
	//"github.com/olekukonko/tablewriter"
)

const bootpath = "boot_target/"

var bootdevicetype = "system"

type boottargetresponse struct {
	Target string `json:"target"`
	status string `json:"status"`
}

func GetCurrentBootTarget(vc *client.VeshClient, rack_id int, board_id int) error {
	status := &boottargetresponse{}
	resp, err := vc.Sling.New().Path(bootpath).Path(strconv.Itoa(rack_id) + "/").Path(strconv.Itoa(board_id) + "/").Get(bootdevicetype).ReceiveSuccess(status)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return err
	}
	fmt.Println(status.Target)
	return nil
}

func SetCurrentBootTarget(vc *client.VeshClient, rack_id int, board_id int, boot_target string) error {
	status := &boottargetresponse{}
	resp, err := vc.Sling.New().Path(bootpath).Path(strconv.Itoa(rack_id) + "/").Path(strconv.Itoa(board_id) + "/").Path(bootdevicetype).Get(boot_target).ReceiveSuccess(status)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return err
	}
	fmt.Println(status.Target, status.status)
	return nil
}
