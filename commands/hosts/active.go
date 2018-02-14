package hosts

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/config"
	"github.com/vapor-ware/synse-cli/flags"
	"github.com/vapor-ware/synse-cli/utils"
)

// hostsActiveCommand is the CLI sub-command for getting the active host.
var hostsActiveCommand = cli.Command{
	Name:  "active",
	Usage: "Display information for the active host",
	Flags: []cli.Flag{
		flags.OutputFlag,
	},
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
	return utils.FormatOutput(c, config.Config.ActiveHost)
}
