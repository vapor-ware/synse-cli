package plugin

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/completion"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-server-grpc/go"
)

const (
	// writeCmdName is the name for the 'write' command.
	writeCmdName = "write"

	// writeCmdUsage is the usage text for the 'write' command.
	writeCmdUsage = "Write data to a plugin's device"

	// writeCmdArgsUsage is the argument usage for the 'write' command.
	writeCmdArgsUsage = "RACK BOARD DEVICE ACTION [DATA]"

	// writeCmdDesc is the description for the 'write' command.
	writeCmdDesc = `The write command writes data to a plugin's device via the
  Synse gRPC API. The plugin write info return is similar to that
  of a 'synse server write' command, and the response data for
  both should look the same.

  Writes to plugins are asynchronous, just as they are for
  Synse Server, so this command will return a transaction ID
  whose state/status can be checked with the 'plugin transaction'
  command.

  When writing to a device, the rack, board, and device id must be
  specified along with the stuff to write. This 'stuff' is composed
  of two parts -- the ACTION (e.g. the thing to change) and the
  DATA (e.g. the value to change it to). Most devices require both
  ACTION and DATA for writing, but some may require only an ACTION.

  Below is a table listing some common actions and the requirements
  for their data

  TYPE      ACTION    DATA
  --------  --------  -------------------
  led       state     (on|off)
            blink     (blink|steady)
            color     RBG HEX string

  fan       speed     integer

Example:
  synse plugin write rack-1 board 29d1a03e8cddfbf1cf68e14e60e5f5cc color ff00ff

Formatting:
  The 'plugin write' command supports the following formatting
  options (via the CLI global --format flag):
    - pretty (default)
		- yaml
    - json`
)

// pluginWriteCommand is a CLI sub-command for writing to a plugin
var pluginWriteCommand = cli.Command{
	Name:        writeCmdName,
	Usage:       writeCmdUsage,
	Description: writeCmdDesc,
	ArgsUsage:   writeCmdArgsUsage,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdWrite(c))
	},

	BashComplete: completion.CompleteRackBoardDevicePlugin,
}

// cmdWrite is the action for pluginWriteCommand. It writes directly to
// the specified plugin.
func cmdWrite(c *cli.Context) error { // nolint: gocyclo
	err := utils.RequiresArgsInRange(4, 5, c)
	if err != nil {
		return err
	}

	rack := c.Args().Get(0)
	board := c.Args().Get(1)
	device := c.Args().Get(2)
	action := c.Args().Get(3)
	raw := c.Args().Get(4)

	wd := &synse.WriteData{
		Action: action,
	}
	if raw != "" {
		wd.Raw = [][]byte{[]byte(raw)}
	}

	transactions, err := client.Grpc.Write(c, rack, board, device, wd)
	if err != nil {
		return err
	}
	t := make([]scheme.WriteTransaction, len(transactions.Transactions))
	for id, ctx := range transactions.Transactions {
		var raw []string
		for _, r := range ctx.Raw {
			raw = append(raw, string(r))
		}

		t = append(t, scheme.WriteTransaction{
			Transaction: id,
			Context: scheme.WriteContext{
				Action: ctx.Action,
				Raw:    raw,
			},
		})
	}

	formatter := formatters.NewWriteFormatter(c, t)
	err = formatter.Add(t)
	if err != nil {
		return err
	}
	return formatter.Write()
}
