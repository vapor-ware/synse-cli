package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/completion"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// readCmdName is the name for the 'read' command.
	readCmdName = "read"

	// readCmdUsage is the usage text for the 'read' command.
	readCmdUsage = "Read from the specified device"

	// readCmdDesc is the description for the 'read' command.
	readCmdDesc = `The read command hits the active Synse Server host's '/read'
	 endpoint to read from the specified device. A Synse Server read
	 will be passed along to the backend plugin which handles the
	 given device to get the reading information. Not all devices
	 may support reading; device support for read is specified at
	 the plugin level.`
)

// ReadCommand is the CLI command for Synse Server's "read" API route.
var ReadCommand = cli.Command{
	Name:        readCmdName,
	Usage:       readCmdUsage,
	Description: readCmdDesc,
	Category:    SynseActionsCategory,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdRead(c))
	},

	BashComplete: completion.CompleteRackBoardDevice,
}

// cmdRead is the action for the ReadCommand. It makes an "read" request
// against the active Synse Server instance.
func cmdRead(c *cli.Context) error {
	err := utils.RequiresArgsExact(3, c)
	if err != nil {
		return err
	}

	rack := c.Args().Get(0)
	board := c.Args().Get(1)
	device := c.Args().Get(2)

	read, err := client.Client.Read(rack, board, device)
	if err != nil {
		return err
	}

	formatter := formatters.NewReadFormatter(c)
	err = formatter.Add(read)
	if err != nil {
		return err
	}
	return formatter.Write()
}
