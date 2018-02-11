package commands

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/config"
	"github.com/vapor-ware/synse-cli/utils"
	"github.com/vapor-ware/synse-cli/flags"
)

// configCommand is the CLI command for displaying the current CLI configuration.
var configCommand = cli.Command{
	Name:   "config",
	Usage:  "Display the current CLI configuration",
	Flags: []cli.Flag{
		flags.OutputFlag,
	},
	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdConfig(c))
	},
}

// cmdConfig is the action for configCommand. It prints out the configuration
// currently set for the CLI.
func cmdConfig(c *cli.Context) error {
	return utils.FormatOutput(c, config.Config)
}
