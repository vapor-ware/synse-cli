package plugin

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/completion"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

// pluginReadCommand is a CLI sub-command for getting a reading from a plugin.
var pluginReadCommand = cli.Command{
	Name:  "read",
	Usage: "Get a reading from a plugin",

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdRead(c))
	},

	BashComplete: completion.CompleteRackBoardDevicePlugin,
}

// cmdRead is the action for pluginReadCommand. It prints out a reading that was
// retrieved from the specified plugin.
func cmdRead(c *cli.Context) error { // nolint: gocyclo
	err := utils.RequiresArgsExact(3, c)
	if err != nil {
		return err
	}

	rack := c.Args().Get(0)
	board := c.Args().Get(1)
	device := c.Args().Get(2)

	resp, err := client.Grpc.Read(c, rack, board, device)
	if err != nil {
		return err
	}

	formatter := formatters.NewReadFormatter(c.App.Writer)
	for _, read := range resp {
		err = formatter.Add(&scheme.Read{
			Type: read.Type,
			Data: map[string]scheme.ReadData{
				read.Type: {
					Value:     read.Value,
					Timestamp: read.Timestamp,
				},
			},
		})
		if err != nil {
			return err
		}
	}
	return formatter.Write()
}
