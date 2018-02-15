package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/completion"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// infoCmdName is the name for the 'info' command.
	infoCmdName = "info"

	// infoCmdUsage is the usage text for the 'info' command.
	infoCmdUsage = "Get info for the specified rack, board, or device"

	// infoCmdDesc is the description for the 'info' command.
	infoCmdDesc = `The info command hits the active Synse Server host's '/info'
	 endpoint. Information can be provided at various scopes: the
	 rack level, the board level, or the device level.`
)

// InfoCommand is the CLI command for Synse Server's "info" API route.
var InfoCommand = cli.Command{
	Name:        infoCmdName,
	Usage:       infoCmdUsage,
	Description: infoCmdDesc,
	Category:    SynseActionsCategory,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdInfo(c))
	},

	BashComplete: completion.CompleteRackBoardDevice,

	Flags: []cli.Flag{
		// --output, -o flag specifies the output format (YAML, JSON) for the command
		cli.StringFlag{
			Name:  "output, o",
			Value: "yaml",
			Usage: "set the output format of the command",
		},
	},
}

// cmdInfo is the action for the InfoCommand. It makes an "info" request
// against the active Synse Server instance.
func cmdInfo(c *cli.Context) (err error) {
	err = utils.RequiresArgsInRange(1, 3, c)
	if err != nil {
		return err
	}

	rack := c.Args().Get(0)
	board := c.Args().Get(1)
	device := c.Args().Get(2)

	var info interface{}

	if board == "" {
		// No board is defined, so we are querying at the rack level.
		info, err = client.Client.RackInfo(rack)

	} else if device == "" {
		// Board is defined, but device is not, so we are querying at the board level.
		info, err = client.Client.BoardInfo(rack, board)

	} else {
		// Rack, Board, Device is defined, so we are querying at the device level.
		info, err = client.Client.DeviceInfo(rack, board, device)
	}

	if err != nil {
		return err
	}

	return formatters.FormatOutput(c, info)
}
