package server

import (
	"fmt"
	"sort"
	"strings"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
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

// TODO (etd) - better organization here. this should probably move to the
// utils or other sorting/filtering package

// Filter is used to filter the scan results based on the given filtering
// function.
func Filter(devices []*scheme.InternalScan, f func(scan *scheme.InternalScan) bool) []*scheme.InternalScan {
	tmp := make([]*scheme.InternalScan, 0)
	for _, d := range devices {
		if f(d) {
			tmp = append(tmp, d)
		}
	}
	return tmp
}

type byScanDeviceID []*scheme.InternalScan

func (s byScanDeviceID) Len() int {
	return len(s)
}

func (s byScanDeviceID) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byScanDeviceID) Less(i, j int) bool {
	return s[i].ID() < s[j].ID()
}

// cmdScan is the action for the ScanCommand. It makes an "scan" request
// against the active Synse Server instance.
func cmdScan(c *cli.Context) error {
	scan, err := client.Client.Scan()
	if err != nil {
		return err
	}

	// convert the scan results to an internal representation that
	// makes it easier to do device-based actions (sorting, filtering,
	// formatting, etc).
	devices := scan.ToInternalScan()

	// Sort by ID
	sort.Sort(byScanDeviceID(devices))

	filterString := c.String("filter")
	if filterString != "" {
		filter := strings.Split(filterString, "=")
		if len(filter) != 2 {
			return cli.NewExitError("filter must be in the form 'key=value'", 1)
		}

		switch strings.ToLower(filter[0]) {
		case "type":
			devices = Filter(devices, func(d *scheme.InternalScan) bool {
				return d.Type == strings.ToLower(filter[1])
			})
		default:
			return cli.NewExitError(
				fmt.Sprintf("filter key for '%v' is not supported", filterString),
				1,
			)
		}
	}

	formatter := formatters.NewScanFormatter(c.App.Writer)
	err = formatter.Add(devices)
	if err != nil {
		return err
	}
	return formatter.Write()
}
