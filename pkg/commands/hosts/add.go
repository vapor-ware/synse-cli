package hosts

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// addCmdName is the name for the 'add' command.
	addCmdName = "add"

	// addCmdUsage is the usage text for the 'add' command.
	addCmdUsage = "Add a new Synse Server host"

	// addCmdArgsUsage is the argument usage for the 'add' command.
	addCmdArgsUsage = "NAME HOST"

	// addCmdDesc is the description for the 'add' command.
	addCmdDesc = `The add command adds a new Synse Server instance to the
  available hosts that the CLI can interface with.

Example:
  synse hosts add local localhost:5000`
)

// hostsAddCommand is the CLI sub-command for adding a new host to the CLI
// configuration.
var hostsAddCommand = cli.Command{
	Name:        addCmdName,
	Usage:       addCmdUsage,
	Description: addCmdDesc,
	ArgsUsage:   addCmdArgsUsage,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdAdd(c))
	},
}

// cmdAdd is the action for hostsAddCommand. It adds the specified host to the
// CLI configuration.
func cmdAdd(c *cli.Context) error {
	err := utils.RequiresArgsExact(2, c)
	if err != nil {
		return err
	}

	name := c.Args().Get(0)
	addr := c.Args().Get(1)

	err = config.Config.AddHost(config.NewHostConfig(name, addr))
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	return nil
}
