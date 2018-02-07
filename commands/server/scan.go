package server

import (
	"fmt"
	"sort"
	"strings"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/flags"
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
	Flags: []cli.Flag{
		flags.FilterFlag,
	},
	Action: func(c *cli.Context) error {
		return utils.CmdHandler(c, cmdScan(c))
	},
}

// scanDevice represents a single device found during a scan.
type scanDevice struct {
	Rack   string
	Board  string
	Device string
	Info   string
	Type   string
}

// ID generates the ID of the device by joining the rack, board, and device.
func (device *scanDevice) ID() string {
	return strings.Join([]string{
		device.Rack,
		device.Board,
		device.Device,
	}, "-")
}

// ToRow converts a scanDevice to a table row.
func (device *scanDevice) ToRow() []string {
	return []string{
		device.ID(),
		device.Info,
		device.Type,
	}
}

// TODO (etd) - better organization here. this should probably move to the
// utils or other sorting/filtering package

func Filter(devices []*scanDevice, f func(*scanDevice) bool) []*scanDevice {
	tmp := make([]*scanDevice, 0)
	for _, d := range devices {
		if f(d) {
			tmp = append(tmp, d)
		}
	}
	return tmp
}

type byScanDeviceID []*scanDevice

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
	scan := &scheme.Scan{}
	err := utils.DoGet(scanBase, scan)
	if err != nil {
		return err
	}

	var devices []*scanDevice
	for _, rack := range scan.Racks {
		for _, board := range rack.Boards {
			for _, device := range board.Devices {
				devices = append(devices, &scanDevice{
					Rack:   rack.ID,
					Board:  board.ID,
					Device: device.ID,
					Info:   device.Info,
					Type:   device.Type,
				})
			}
		}
	}

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
			devices = Filter(devices, func(d *scanDevice) bool {
				return d.Type == strings.ToLower(filter[1])
			})
		default:
			return cli.NewExitError(
				fmt.Sprintf("filter key for '%v' is not supported", filterString),
				1,
			)
		}
	}

	var data [][]string
	for _, dev := range devices {
		data = append(data, dev.ToRow())
	}

	header := []string{"ID", "Info", "Type"}
	utils.TableOutput(header, data)
	return nil
}
