package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
)

// infoBase is the base URI for the "info" route.
const infoBase = "info"

// InfoCommand is the CLI command for Synse Server's "info" API route.
var InfoCommand = cli.Command{
	Name:     "info",
	Usage:    "Get info for the specified rack, board, or device",
	Category: "Synse Server Actions",
	Action: func(c *cli.Context) error {
		return utils.CmdHandler(c, cmdInfo(c))
	},
}

// cmdInfo is the action for the InfoCommand. It makes an "info" request
// against the active Synse Server instance.
func cmdInfo(c *cli.Context) error {
	err := utils.RequiresArgsInRange(1, 3, c)
	if err != nil {
		return err
	}

	rack := c.Args().Get(0)
	board := c.Args().Get(1)
	device := c.Args().Get(2)

	var info interface{}
	var uri string

	if board == "" {
		// No board is defined, so we are querying at the rack level.
		info = &scheme.RackInfo{}
		uri = utils.MakeURI(infoBase, rack)

	} else if device == "" {
		// Board is defined, but device is not, so we are querying at the board level.
		info = &scheme.BoardInfo{}
		uri = utils.MakeURI(infoBase, rack, board)

	} else {
		// Rack, Board, Device is defined, so we are querying at the device level.
		info = &scheme.DeviceInfo{}
		uri = utils.MakeURI(infoBase, rack, board, device)
	}

	err = utils.DoGet(uri, info)
	if err != nil {
		return err
	}

	return utils.AsYAML(info)
}
