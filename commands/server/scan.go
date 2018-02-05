package server

import (
	"net/http"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
)

// scanBase is the base URI for the "scan" route.
const scanBase = "scan"

// ScanCommand is the CLI command for Synse Server's "scan" API route.
var ScanCommand = cli.Command{
	Name:     "scan",
	Usage:    "Enumerate all devices on the active host",
	Category: "Synse Server Actions",
	Action: func(c *cli.Context) error {
		return utils.CommandHandler(c, cmdScan(c))
	},
}

// cmdScan is the action for the ScanCommand. It makes an "scan" request
// against the active Synse Server instance.
func cmdScan(c *cli.Context) error {
	scan := &scheme.Scan{}
	resp, err := client.New().Get(scanBase).ReceiveSuccess(scan)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return err
	}

	var data [][]string
	for _, rack := range scan.Racks {
		for _, board := range rack.Boards {
			for _, device := range board.Devices {
				data = append(data, []string{
					rack.Id,
					board.Id,
					device.Id,
					device.Info,
					device.Type,
				})
			}
		}
	}

	// FIXME (etd) - this should be sorted

	header := []string{"Rack", "Board", "Device", "Info", "Type"}
	utils.TableOutput(header, data)
	return nil
}
