package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/completion"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// writeCmdName is the name for the 'write' command.
	writeCmdName = "write"

	// writeCmdUsage is the usage text for the 'write' command.
	writeCmdUsage = "Write to the specified device"

	// writeCmdArgsUsage is the argument usage for the 'write' command.
	writeCmdArgsUsage = "RACK BOARD DEVICE ACTION [DATA]"

	// writeCmdDesc is the description for the 'write' command.
	writeCmdDesc = `The write command hits the active Synse Server host's '/write'
  endpoint to write to the specified device. A Synse Server write
  will be passed along to the backend plugin which handles the
  given device to get the write information. Not all devices
  may support writing; support for write is determined at the
  plugin level.

  Writes are asynchronous, so issuing a write command will return
  a transaction ID. This ID can be checked to get the state/status
  of the write using the 'synse server transaction' command.

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
  synse server write rack-1 board 29d1a03e8cddfbf1cf68e14e60e5f5cc color ff00ff

Formatting:
  The 'server write' command supports the following formatting
  options (via the CLI global --format flag):
    - pretty (default)
		- yaml
    - json`
)

// writeCommand is the CLI command for Synse Server's "write" API route.
var writeCommand = cli.Command{
	Name:        writeCmdName,
	Usage:       writeCmdUsage,
	Description: writeCmdDesc,
	ArgsUsage:   writeCmdArgsUsage,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdWrite(c))
	},

	BashComplete: completion.CompleteRackBoardDevice,
}

// cmdWrite is the action for the writeCommand. It makes a "write" request
// against the active Synse Server instance.
func cmdWrite(c *cli.Context) error {
	err := utils.RequiresArgsInRange(4, 5, c)
	if err != nil {
		return err
	}

	rack := c.Args().Get(0)
	board := c.Args().Get(1)
	device := c.Args().Get(2)
	action := c.Args().Get(3)
	raw := c.Args().Get(4)

	write, err := client.Client.Write(rack, board, device, action, raw)
	if err != nil {
		return err
	}

	formatter := formatters.NewWriteFormatter(c, write)
	err = formatter.Add(write)
	if err != nil {
		return err
	}
	return formatter.Write()
}
