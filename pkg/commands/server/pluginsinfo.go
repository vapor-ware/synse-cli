package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// pluginsInfoCmdName is the name for the 'plugins info' command.
	pluginsInfoCmdName = "info"

	// pluginsInfoCmdUsage is the usage text for the 'plugins info' command.
	pluginsInfoCmdUsage = "Get the list of plugins' metadata that are configured with Synse Server"

	// pluginsInfoCmdDesc is the description for the 'plugins info' command.
	pluginsInfoCmdDesc = `The plugins info command hits the active Synse Server host's '/plugins'
  endpoint, returns metadata information of all configured plugins.

Example:
  synse server plugins info

Formatting:
  The 'server plugins info' command supports the following formatting
  options (via the CLI global --format flag):
    - yaml (default)
    - json`
)

// pluginsInfoCommand is additional CLI command for Synse Server's "plugins" API route.
var pluginsInfoCommand = cli.Command{
	Name:        pluginsInfoCmdName,
	Usage:       pluginsInfoCmdUsage,
	Description: pluginsInfoCmdDesc,
	ArgsUsage:   utils.NoArgs,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdPluginsInfo(c))
	},
}

// cmPluginsInfo is the action for the pluginsInfoCommand. It makes a "plugins" request
// against the active Synse Server instance and returns plugins' metadata information.
func cmdPluginsInfo(c *cli.Context) error {
	plugins, err := client.Client.Plugins()
	if err != nil {
		return err
	}

	formatter := formatters.NewServerPluginsInfoFormatter(c)
	err = formatter.Add(plugins)
	if err != nil {
		return err
	}
	return formatter.Write()
}
