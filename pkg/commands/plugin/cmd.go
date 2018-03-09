package plugin

import (
	"github.com/urfave/cli"
)

const (
	cmdName = "plugin"

	cmdUsage = "Interact with Synse Plugins"

	cmdDescription = `This sub-command allows you to interact with any Synse
  Plugin instance. It uses the Synse gRPC API to make requests
  directly to the Plugin itself. This does not depend on Synse
  Server, so it can be useful for debugging plugins themselves.

  One of the '--tcp' or '--unix' flags must be specified for
  the CLI to know how to interface with the Plugin.`
)

// PluginCommand is the CLI command for interacting with Synse plugins.
var PluginCommand = cli.Command{
	Name:        cmdName,
	Usage:       cmdUsage,
	Description: cmdDescription,

	Subcommands: []cli.Command{
		pluginMetainfoCommand,
		pluginReadCommand,
		pluginTransactionCommand,
		pluginWriteCommand,
	},

	Flags: []cli.Flag{
		// --tcp, -t flag specifies the TCP address for the plugin to interface with
		cli.StringFlag{
			Name:  "tcp, t",
			Usage: "set the hostname/ip[:port] for a plugin",
		},
		// --unix, -u flag specifies the UNIX socket for the plugin to interface with
		cli.StringFlag{
			Name:  "unix, u",
			Usage: "set the unix socket path for a plugin",
		},
	},
}
