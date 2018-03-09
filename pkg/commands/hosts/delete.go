package hosts

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/completion"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// deleteCmdName is the name for the 'delete' command.
	deleteCmdName = "delete"

	// deleteCmdUsage is the usage text for the 'delete' command.
	deleteCmdUsage = "Delete a Synse Server host"

	// deleteCmdArgsUsage is the argument usage for the 'delete' command.
	deleteCmdArgsUsage = "NAME"

	// deleteCmdDesc is the description for the 'delete' command.
	deleteCmdDesc = `The delete command removes a Synse Server instance from the
  tracked hosts configuration.

Example:
  synse hosts delete local`
)

// hostsDeleteCommand is the CLI sub-command for deleting a host from the
// CLI configuration.
var hostsDeleteCommand = cli.Command{
	Name:        deleteCmdName,
	Usage:       deleteCmdUsage,
	Description: deleteCmdDesc,
	ArgsUsage:   deleteCmdArgsUsage,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdDelete(c))
	},

	BashComplete: completion.CompleteHostNames,
}

// cmdDelete is the action for hostsDeleteCommand. It removes the specified host
// from the CLI configuration, if it exists. If the specified host is also the
// active host, it will unset the active host.
func cmdDelete(c *cli.Context) error {
	err := utils.RequiresArgsExact(1, c)
	if err != nil {
		return err
	}

	name := c.Args().Get(0)

	host := config.Config.Hosts[name]
	if host != nil {
		if config.Config.ActiveHost != nil && *host == *config.Config.ActiveHost {
			config.Config.ActiveHost = nil
		}
	}
	delete(config.Config.Hosts, name)
	return nil
}
