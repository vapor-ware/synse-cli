package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
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
		cli.StringFlag{
			Name:  "filter, f",
			Usage: "set a filter for the output results",
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

	transformer := scan.ToScanTransformer()
	transformer.OrderBy("rack", "board", "device")
	transformer.FiltersFromContext(c)
	transformer.Sort()
	transformer.Filter()

	formatter := formatters.NewScanFormatter(c)
	err = formatter.Add(transformer.Items)
	if err != nil {
		return err
	}
	return formatter.Write()
}
