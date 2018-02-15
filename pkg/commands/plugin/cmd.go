package plugin

import (
	"github.com/urfave/cli"
)

// PluginCommand is the CLI command for interacting with Synse plugins.
var PluginCommand = cli.Command{
	Name:  "plugin",
	Usage: "Interact with Synse Plugins",

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
