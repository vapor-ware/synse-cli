package server

import (
	"net/http"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
)

// scanURI
const scanURI = "scan"

// statusCommand
var ScanCommand = cli.Command{
	Name:  "scan",
	Usage: "scan the Synse Server instance",
	Action: func(c *cli.Context) error {
		return utils.CommandHandler(c, cmdScan(c))
	},
}

// cmdScan
func cmdScan(c *cli.Context) error {
	scan := &scheme.Scan{}
	resp, err := client.New().Get(scanURI).ReceiveSuccess(scan)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return err
	}

	// FIXME -- the below doesn't work anymore ):
	var data [][]string

	filter := &utils.FilterFunc{}
	filter.FilterFn = func(res utils.Result) bool {
		return true
	}

	fil, err := utils.FilterDevices(filter)
	if err != nil {
		return err
	}
	for res := range fil {
		if res.Error != nil {
			return res.Error
		}
		data = append(data, []string{})
	}

	header := []string{"Rack", "Board", "Device", "Info", "Type"}
	utils.TableOutput(header, data)
	return nil
}
