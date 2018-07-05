package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// pluginsInfoCmdName is the name for the 'plugins info' command.
	pluginsInfoCmdName = "info"

	// pluginsInfoCmdUsage is the usage text for the 'plugins info' command.
	pluginsInfoCmdUsage = "Get a list of plugins' metadata that are configured with Synse Server"

	// pluginsInfoCmdArgsUsage is the argument usage for the `plugins info`
	// command.
	pluginsInfoCmdArgsUsage = "[PLUGIN TAG]"

	// pluginsInfoCmdDesc is the description for the 'plugins info' command.
	pluginsInfoCmdDesc = `The plugins info command hits the active Synse Server host's '/plugins'
  endpoint. If a plugin is provided, the CLI will return its metadata
  information. Otherwise, it returns metadata information of all
  configured plugins.

Example:
  # Get metadata of all configured plugins (default)
  synse server plugins info

  # Get metadata of vaporio/emulator-plugin
  synse server plugins info emulator-plugin

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

	plugins, err := getPlugins(c.Args().Get(0), c)
	if err != nil {
		return err
	}

	if len(plugins) == 0 {
		return nil
	}

	formatter := formatters.NewServerPluginsInfoFormatter(c)
	err = formatter.Add(plugins)
	if err != nil {
		return err
	}
	return formatter.Write()
}

// getPlugins is a helper function that takes the given plugin tag and returns
// the set of matched plugins.
func getPlugins(pluginTag string, c *cli.Context) ([]scheme.Plugin, error) {
	var plugins []scheme.Plugin

	pluginsResults, err := client.Client.Plugins()
	if err != nil {
		return nil, err
	}

	if pluginTag == "" {
		return pluginsResults, nil
	}

	for _, plugin := range pluginsResults {
		if pluginTag == plugin.Tag {
			plugins = append(plugins, plugin)
		}
	}

	return plugins, nil
}
