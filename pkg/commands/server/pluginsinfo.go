package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// pluginsInfoCmdName is the name for the 'plugins info' command.
	pluginsInfoCmdName = "info"

	// pluginsInfoCmdUsage is the usage text for the 'plugins info' command.
	pluginsInfoCmdUsage = "Get a list of plugins' metadata that are configured with Synse Server"

	// pluginsInfoCmdArgsUsage is the argument usage for the `plugins info` command.
	pluginsInfoCmdArgsUsage = "[PLUGIN TAG ...]"

	// pluginsInfoCmdDesc is the description for the 'plugins info' command.
	pluginsInfoCmdDesc = `The plugins info command hits the active Synse Server host's '/plugins'
  endpoint. If a plugin tag or mutiple plugins' tags (up to 3) are provided,
  the CLI returns their metadata information. Otherwise, it returns metadata
  information of all configured plugins.

Example:
  # Get metadata of all configured plugins (default)
  synse server plugins info

  # Get metadata of vaporio/emulator-plugin
  synse server plugins info vaporio/emulator-plugin

Formatting:
  The 'server plugins info' command supports the following formatting
  options (via the CLI global --format flag):
    - yaml (default)
    - json`
)

// pluginsInfoCommand is a CLI command for Synse Server's "plugins" API route
// that gets metadata information.
var pluginsInfoCommand = cli.Command{
	Name:        pluginsInfoCmdName,
	Usage:       pluginsInfoCmdUsage,
	Description: pluginsInfoCmdDesc,
	ArgsUsage:   pluginsInfoCmdArgsUsage,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdPluginsInfo(c))
	},
}

// cmPluginsInfo is the action for the pluginsInfoCommand. It makes a "plugins" request
// against the active Synse Server instance and returns plugins' metadata information.
func cmdPluginsInfo(c *cli.Context) error {
	err := utils.RequiresArgsInRange(0, 1, c)
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

	formatter := formatters.NewServerPluginsInfoFormatter(c)
	err = formatter.Add(plugins)
	if err != nil {
		return err
	}
	return formatter.Write()
}
