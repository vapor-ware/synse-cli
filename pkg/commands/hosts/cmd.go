package hosts

import (
	"github.com/urfave/cli"
)

const (
	cmdName = "hosts"

	cmdUsage = "Manage Synse Server instances"

	cmdDescription = `This sub-command allows you to add, remove, change, and
  enumerate Synse Server instance that can be interfaced with
  for 'synse server' commands.`
)

// HostsCommand is the CLI command for managing Synse Server hosts.
var HostsCommand = cli.Command{
	Name:        cmdName,
	Usage:       cmdUsage,
	Description: cmdDescription,

	Subcommands: []cli.Command{
		hostsActiveCommand,
		hostsAddCommand,
		hostsChangeCommand,
		hostsDeleteCommand,
		hostsListCommand,
	},
}
