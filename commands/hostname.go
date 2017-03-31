package commands

import (
	"strings"

	"github.com/vapor-ware/vesh/client"
	"github.com/vapor-ware/vesh/utils"
)

// FIXME: Break printing back out into another function
// ListHostnames iterates over the complete list of boards and returns the
// hostname(s) and ip address(es) associated with each, given from the top
// level "hostnames" and "ip addresses" fields. Since a given board may have
// multiple hostnames and/or ip addresses, all given values for each field are
// returned.
func ListHostnames(vc *client.VeshClient) error {
	var data [][]string

	filter := &utils.FilterFunc{}
	filter.FilterFn = func(res utils.Result) bool {
		return res.DeviceType == "system"
	}

	fil, err := utils.FilterDevices(filter)
	if err != nil {
		return err
	}
	for res := range fil {
		if res.Error != nil {
			return res.Error
		}
		data = append(data, []string{
			res.RackID,
			res.BoardID,
			strings.Join(res.Hostnames, ","),
			strings.Join(res.IPAddresses, ",")})
	}

	header := []string{"Rack", "Board", "Hostnames", "IP Addesses"}
	utils.TableOutput(header, data)

	return err
}

// FIXME: Break getting back out into another function
// PrintGetHostname takes the output of GetHostname and pretty prints it in table form.
func PrintGetHostname(vc *client.VeshClient, rack_id, board_id string) error {
	var data [][]string

	filter := &utils.FilterFunc{}
	filter.FilterFn = func(res utils.Result) bool {
		return res.DeviceType == "system" && res.RackID == rack_id && res.BoardID == board_id
	}

	fil, err := utils.FilterDevices(filter)
	if err != nil {
		return err
	}
	for res := range fil {
		if res.Error != nil {
			return res.Error
		}
		data = append(data, []string{
			strings.Join(res.Hostnames, ","),
			strings.Join(res.IPAddresses, ",")})
	}

	header := []string{"Hostnames", "IP Addesses"}
	utils.TableOutput(header, data)

	return err
}
