package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
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

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdConfig(c))
	},
}

// cmdConfig is the action for the ConfigCommand. It makes an "config" request
// against the active Synse Server instance.
func cmdConfig(c *cli.Context) error {
	cfg, err := client.Client.Config()
	if err != nil {
		return err
	}

	formatter := formatters.NewConfigFormatter(c, cfg)
	return formatter.Write()
}
