package hosts

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/completion"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// changeCmdName is the name for the 'change' command.
	changeCmdName = "change"

	// changeCmdUsage is the usage text for the 'change' command.
	changeCmdUsage = "Change the active host"

	// changeCmdArgsUsage is the argument usage for the 'change' command.
	changemdArgsUsage = "NAME"

	// changeCmdDesc is the description for the 'change' command.
	changeCmdDesc = `The change command changes the active host to the one specified
  by the given reference name. To see a list of all hosts that
  are currently configured with the CLI, use 'synse hosts list'.

Example:
  synse hosts change local`
)

// hostsChangeCommand is the CLI sub-command for changing the active host.
var hostsChangeCommand = cli.Command{
	Name:        changeCmdName,
	Usage:       changeCmdUsage,
	Description: changeCmdDesc,
	ArgsUsage:   changemdArgsUsage,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdChange(c))
	},

	BashComplete: completion.CompleteHostNames,
}

// cmdChange is the action for hostsChangeCommand. It changes the active host to
// the specified host, if it exists.
func cmdChange(c *cli.Context) error {
	err := utils.RequiresArgsExact(1, c)
	if err != nil {
		return err
	}

	name := c.Args().Get(0)

	for _, host := range config.Config.Hosts {
		if host.Name == name {
			config.Config.ActiveHost = host
			return nil
		}
	}
	return cli.NewExitError(fmt.Sprintf("host with name '%v' not found", name), 1)
}
