package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/completion"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// infoCmdName is the name for the 'info' command.
	infoCmdName = "info"

	// infoCmdUsage is the usage text for the 'info' command.
	infoCmdUsage = "Get info for the specified rack, board, or device"

	// infoCmdArgsUsage is the argument usage for the 'info' command.
	infoCmdArgsUsage = "RACK [BOARD [DEVICE]]"

	// infoCmdDesc is the description for the 'info' command.
	infoCmdDesc = `The info command hits the active Synse Server host's '/info'
  endpoint. Information can be provided at various scopes: the
  rack level, the board level, or the device level. Each level
  of info will provide locational information (e.g. what lives
  on it/what it lives on) as well as any entity-specific info.

Example:
  # rack level info request
  synse rack-1

  # board level info request
  synse rack-1 board

  # device level info request
  synse rack-1 board 29d1a03e8cddfbf1cf68e14e60e5f5cc

Formatting:
  The 'server info' command supports the following formatting
  options (via the CLI global --format flag):
    - yaml (default)
    - json`
)

// infoCommand is the CLI command for Synse Server's "info" API route.
var infoCommand = cli.Command{
	Name:        infoCmdName,
	Usage:       infoCmdUsage,
	Description: infoCmdDesc,
	ArgsUsage:   infoCmdArgsUsage,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdInfo(c))
	},

	BashComplete: completion.CompleteRackBoardDevice,
}

// cmdInfo is the action for the infoCommand. It makes an "info" request
// against the active Synse Server instance.
func cmdInfo(c *cli.Context) (err error) {
	err = utils.RequiresArgsInRange(1, 3, c)
	if err != nil {
		return err
	}

	rack := c.Args().Get(0)
	board := c.Args().Get(1)
	device := c.Args().Get(2)

	formatter := formatters.NewInfoFormatter(c)

	var info interface{}

	if board == "" {
		// No board is defined, so we are querying at the rack level.
		info, err = client.Client.RackInfo(rack)
		formatter.Decoder = &scheme.RackInfo{}

	} else if device == "" {
		// Board is defined, but device is not, so we are querying at the board level.
		info, err = client.Client.BoardInfo(rack, board)
		formatter.Decoder = &scheme.BoardInfo{}

	} else {
		// Rack, Board, Device is defined, so we are querying at the device level.
		info, err = client.Client.DeviceInfo(rack, board, device)
		formatter.Decoder = &scheme.DeviceInfo{}
	}

	if err != nil {
		return err
	}

	err = formatter.Add(info)
	if err != nil {
		return err
	}

	return formatter.Write()
}
