package hosts

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

// hostsActiveCommand is the CLI sub-command for getting the active host.
var hostsActiveCommand = cli.Command{
	Name:  "active",
	Usage: "Display information for the active host",

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdActive(c))
	},

	Flags: []cli.Flag{
		// --output, -o flag specifies the output format (YAML, JSON) for the command
		cli.StringFlag{
			Name:  "output, o",
			Value: "yaml",
			Usage: "set the output format of the command",
		},
	},
}

// cmdActive is the action for hostsActiveCommand. It prints out the information
// associated with the active host, if one is set.
func cmdActive(c *cli.Context) error {
	if config.Config.ActiveHost == nil {
		return cli.NewExitError("no active host set", 1)
	}
	return formatters.FormatOutput(c, config.Config.ActiveHost)
}
