package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/transformers"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// scanCmdName is the name for the 'scan' command.
	scanCmdName = "scan"

	// scanCmdUsage is the usage text for the 'scan' command.
	scanCmdUsage = "Enumerate all devices on the active host"

	// scanCmdDesc is the description for the 'scan' command.
	scanCmdDesc = `The scan command hits the active Synse Server host's '/scan'
	 endpoint, which enumerates all devices that are known to Synse
	 Server via the instance's configured plugins.`
)

// ScanCommand is the CLI command for Synse Server's "scan" API route.
var ScanCommand = cli.Command{
	Name:        scanCmdName,
	Usage:       scanCmdUsage,
	Description: scanCmdDesc,
	Category:    SynseActionsCategory,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdScan(c))
	},

	Flags: []cli.Flag{
		// --filter, -f flag sets a filter in the form "key=value" on the command output
		cli.StringSliceFlag{
			Name:  "filter, f",
			Usage: "set filters for the output results",
		},
		// --sort flag sets the sorting constraints. multiple constraints should be
		// separated by commas, e.g. "key1,key2".
		cli.StringFlag{
			Name:  "sort",
			Value: "rack,board,device",
			Usage: "set the sort constraints for pretty output",
		},
	},
}

// cmdScan is the action for the ScanCommand. It makes a "scan" request
// against the active Synse Server instance.
func cmdScan(c *cli.Context) error {
	scan, err := client.Client.Scan()
	if err != nil {
		return err
	}

	devices := scan.ToScanDevices()

	t, err := transformers.NewScanTransformer(devices)
	if err != nil {
		return err
	}
	err = t.SetFilters(c)
	if err != nil {
		return err
	}
	t.OrderBy(c.String("sort"))
	t.Apply()

	formatter := formatters.NewScanFormatter(c)
	err = formatter.Add(t.Items)
	if err != nil {
		return err
	}
	return formatter.Write()
}
