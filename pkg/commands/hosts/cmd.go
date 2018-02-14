package hosts

import (
	"github.com/urfave/cli"
)

// HostsCommand is the CLI command for managing Synse Server hosts.
var HostsCommand = cli.Command{
	Name:  "hosts",
	Usage: "Manage Synse Server instances",

	Subcommands: []cli.Command{
		hostsActiveCommand,
		hostsAddCommand,
		hostsChangeCommand,
		hostsDeleteCommand,
		hostsListCommand,
	},
}
