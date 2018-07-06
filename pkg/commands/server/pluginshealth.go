package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// pluginsHealthCmdName is the name for the 'plugins health' command.
	pluginsHealthCmdName = "health"

	// pluginsHealthCmdUsage is the usage text for the 'plugins health' command.
	pluginsHealthCmdUsage = "Get a list of plugins' health that are configured with Synse Server"

	// pluginsHealthCmdArgsUsage is the argument usage for the `plugins health` command.
	pluginsHealthCmdArgsUsage = "[PLUGIN TAG ...]"

	// pluginsHealthCmdDesc is the description for the 'plugins health' command.
	pluginsHealthCmdDesc = `The plugins health command hits the active Synse Server host's '/plugins'
  endpoint. If a plugin tag or mutiple plugins' tags (up to 3) are provided,
  the CLI returns their health information. Otherwise, it returns health 
  information of all configured plugins.

Example:
  # Get health of all configured plugins (default)
  synse server plugins health

  # Get health of vaporio/emulator-plugin
  synse server plugins health vaporio/emulator-plugin

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
	ArgsUsage:   pluginsHealthCmdArgsUsage,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdPluginsHealth(c))
	},
}

// cmPluginsHealth is the action for the pluginsHealthCommand. It makes a "plugins" request
// against the active Synse Server instance and returns plugins' health information.
func cmdPluginsHealth(c *cli.Context) error {
	err := utils.RequiresArgsInRange(0, 3, c)
	if err != nil {
		return err
	}

	plugins, err := getPlugins(
		c,
		c.Args().Get(0),
		c.Args().Get(1),
		c.Args().Get(2),
	)
	if err != nil {
		return err
	}

	// FIXME: If plugins is empty, formatter raises a "no data to write error".
	// Refer to #187's comment.

	formatter := formatters.NewServerPluginsHealthFormatter(c)
	err = formatter.Add(plugins)
	if err != nil {
		return err
	}
	return formatter.Write()
}
