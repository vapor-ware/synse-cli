package hosts

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// activeCmdName is the name for the 'active' command.
	activeCmdName = "active"

	// activeCmdUsage is the usage text for the 'active' command.
	activeCmdUsage = "Show the active host"

	// activeCmdDesc is the description for the 'active' command.
	activeCmdDesc = `The active command shows information on which Synse Server
  host is currently active. An 'active host' is the instance
  that all 'synse server' commands will be routed to.

Example:
  synse hosts active

Formatting:
  The 'hosts active' command supports the following formatting
  options (via the CLI global --format flag):
    - yaml (default)
    - json`
)

// hostsActiveCommand is the CLI sub-command for getting the active host.
var hostsActiveCommand = cli.Command{
	Name:        activeCmdName,
	Usage:       activeCmdUsage,
	Description: activeCmdDesc,
	ArgsUsage:   utils.NoArgs,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdActive(c))
	},
}

// cmdActive is the action for hostsActiveCommand. It prints out the information
// associated with the active host, if one is set.
func cmdActive(c *cli.Context) error {
	if config.Config.ActiveHost == nil {
		return cli.NewExitError("no active host set", 1)
	}

	formatter := formatters.NewActiveFormatter(c)
	err := formatter.Add(config.Config.ActiveHost)
	if err != nil {
		return err
	}
	return formatter.Write()
}
