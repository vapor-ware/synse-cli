package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// pluginsCmdName is the name for the 'plugins' command.
	pluginsCmdName = "plugins"

	// pluginsCmdUsage is the usage text for the 'plugins' command.
	pluginsCmdUsage = "Get the list of plugins that are configured with Synse Server"

	// pluginsCmdDesc is the description for the 'plugins' command.
	pluginsCmdDesc = `The plugins command hits the active Synse Server host's '/plugins'
  endpoint, which returns the current set of configured plugins for
  that instance.

Example:
  synse server plugins

Formatting:
  The 'server plugins' command supports the following formatting
  options (via the CLI global --format flag):
    - pretty (default)
    - yaml
    - json`
)

// pluginsCommand is the CLI command for Synse Server's "plugins" API route.
var pluginsCommand = cli.Command{
	Name:        pluginsCmdName,
	Usage:       pluginsCmdUsage,
	Description: pluginsCmdDesc,
	ArgsUsage:   utils.NoArgs,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdPlugins(c))
	},
}

// cmPlugins is the action for the pluginsCommand. It makes a "plugins" request
// against the active Synse Server instance.
func cmdPlugins(c *cli.Context) error {
	plugins, err := client.Client.Plugins()
	if err != nil {
		return err
	}

	formatter := formatters.NewPluginsFormatter(c)
	err = formatter.Add(plugins)
	if err != nil {
		return err
	}
	return formatter.Write()
}
