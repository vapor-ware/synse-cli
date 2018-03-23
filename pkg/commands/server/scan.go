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
  Server via the instance's configured plugins.

  The scan results can be sorted and filtered, which can be
  useful when there are many devices managed by Synse Server.

  Sorting is done via the '--sort' flag. The value for the flag
  specified the fields that should be sorted. Multiple fields
  can be specified by a comma separated string, e.g. "rack,board".
  The first field will be the primary sort key, the second field
  is the secondary sort key, etc. Currently, the fields that are
  supported for sorting are:
    * rack        * device
    * board       * type

  Filtering is done via the '--filter' flag. The value for the
  flag is a string in the format "KEY=VALUE", where KEY is the
  field to filter by, and VALUE is the desired value for that
  field. Multiple filters can be set; see below for an example.
  Currently, the fields that are supported for filtering are:
    * rack        * type
    * board

Example:
  # perform a scan and show only LED devices sorted by their
  # rack, board, and device ids
  synse server scan --sort "rack,board,device" --filter "type=led"

  # perform a scan and show only temperature and pressure
  # devices
  synse server scan --filter "type=temperature" --filter "type=pressure"

Formatting:
  The 'server scan' command supports the following formatting
  options (via the CLI global --format flag):
    - pretty (default)
		- yaml
    - json`
)

// scanCommand is the CLI command for Synse Server's "scan" API route.
var scanCommand = cli.Command{
	Name:        scanCmdName,
	Usage:       scanCmdUsage,
	Description: scanCmdDesc,
	ArgsUsage:   utils.NoArgs,

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

// cmdScan is the action for the scanCommand. It makes a "scan" request
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

	formatter := formatters.NewScanFormatter(c, scan)
	err = formatter.Add(t.Items)
	if err != nil {
		return err
	}
	return formatter.Write()
}
