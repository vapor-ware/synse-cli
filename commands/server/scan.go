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
	Name:     "scan",
	Usage:    "scan the Synse Server instance",
	Category: "Synse Server Actions",
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
