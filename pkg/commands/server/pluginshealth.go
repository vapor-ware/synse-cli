package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// pluginsHealthCmdName is the name for the 'plugins health' command.
	pluginsHealthCmdName = "health"

	// pluginsHealthCmdUsage is the usage text for the 'plugins health' command.
	pluginsHealthCmdUsage = "Get a list of plugins' health that are configured with Synse Server"

	// pluginsHealthCmdDesc is the description for the 'plugins health' command.
	pluginsHealthCmdDesc = `The plugins health command hits the active Synse Server host's '/plugins'
  endpoint, returns health information of all configured plugins.

Example:
  synse server plugins health

Formatting:
  The 'server plugins health' command supports the following formatting
  options (via the CLI global --format flag):
    - yaml (default)
    - json`
)

// pluginsHealthCommand is a CLI command for Synse Server's "plugins" API route
// that gets health information.
var pluginsHealthCommand = cli.Command{
	Name:        pluginsHealthCmdName,
	Usage:       pluginsHealthCmdUsage,
	Description: pluginsHealthCmdDesc,
	ArgsUsage:   utils.NoArgs,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdPluginsHealth(c))
	},
}

// cmPluginsHealth is the action for the pluginsHealthCommand. It makes a "plugins" request
// against the active Synse Server instance and returns plugins' health information.
func cmdPluginsHealth(c *cli.Context) error {
	plugins, err := client.Client.Plugins()
	if err != nil {
		return err
	}

	formatter := formatters.NewServerPluginsHealthFormatter(c)
	err = formatter.Add(plugins)
	if err != nil {
		return err
	}
	return formatter.Write()
}
