package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/flags"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
)

const (
	// configBase is the base URI for the 'config' route.
	configBase = "config"

	// configCmdName is the name for the 'config' command.
	configCmdName = "config"

	// configCmdUsage is the usage text for the 'config' command.
	configCmdUsage = "Get the configuration for the active host"

	// configCmdDesc is the description for the 'config' command.
	configCmdDesc = `The config command hits the active Synse Server host's '/config'
	 endpoint, which returns the current active configuration for that
	 instance.`
)

// ConfigCommand is the CLI command for Synse Server's "config" API route.
var ConfigCommand = cli.Command{
	Name:        configCmdName,
	Usage:       configCmdUsage,
	Description: configCmdDesc,
	Category:    SynseActionsCategory,
	ArgsUsage:   utils.NoArgs,
	Flags: []cli.Flag{
		flags.OutputFlag,
	},
	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdConfig(c))
	},
}

// cmdConfig is the action for the ConfigCommand. It makes an "config" request
// against the active Synse Server instance.
func cmdConfig(c *cli.Context) error {
	cfg := &scheme.Config{}
	err := client.DoGet(configBase, cfg)
	if err != nil {
		return err
	}

	return utils.FormatOutput(c, cfg)
}
