package server

import (
	"sort"
	"strings"

	"github.com/urfave/cli"
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
		return utils.CmdHandler(c, cmdScan(c))
	},
}

type scanDevice struct {
	Rack   string
	Board  string
	Device string
	Info   string
	Type   string
}

// ID
func (device *scanDevice) ID() string {
	return strings.Join([]string{
		device.Rack,
		device.Board,
		device.Device,
	}, "-")
}

// ToRow
func (device *scanDevice) ToRow() []string {
	return []string{
		device.ID(),
		device.Info,
		device.Type,
	}
}

// TODO (etd) - better organization here. this should probably move to the
// utils or other sorting/filtering package
type byScanDeviceId []*scanDevice

func (s byScanDeviceId) Len() int {
	return len(s)
}

func (s byScanDeviceId) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byScanDeviceId) Less(i, j int) bool {
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
					Rack:   rack.Id,
					Board:  board.Id,
					Device: device.Id,
					Info:   device.Info,
					Type:   device.Type,
				})
			}
		}
	}

	// Sort by ID
	sort.Sort(byScanDeviceId(devices))

	var data [][]string
	for _, dev := range devices {
		data = append(data, dev.ToRow())
	}

	header := []string{"ID", "Info", "Type"}
	utils.TableOutput(header, data)
	return nil
}
