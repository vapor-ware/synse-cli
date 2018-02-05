package hosts

import (
	"github.com/urfave/cli"
)

// NewHostsCommand
func NewHostsCommand() cli.Command {
	return cli.Command{
		Name:  "hosts",
		Usage: "manage the configured Synse Server instances",
		Subcommands: []cli.Command{
			hostsActiveCommand,
			hostAddCommand,
			hostChangeCommand,
			hostDeleteCommand,
			hostsListCommand,
		},
	}
}
