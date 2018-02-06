package hosts

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/config"
)

// hostsDeleteCommand is the CLI sub-command for deleting a host from the
// CLI configuration.
var hostsDeleteCommand = cli.Command{
	Name:   "delete",
	Usage:  "Delete a Synse Server host from configuration",
	Action: cmdDelete,
}

// cmdDelete is the action for hostsDeleteCommand. It removes the specified host
// from the CLI configuration, if it exists. If the specified host is also the
// active host, it will unset the active host.
func cmdDelete(c *cli.Context) error {
	name := c.Args().Get(0)
	if name == "" {
		return cli.NewExitError("'delete' requires 1 argument", 1)
	}

	host := config.Config.Hosts[name]
	if host != nil {
		if *host == *config.Config.ActiveHost {
			config.Config.ActiveHost = nil
		}
	}
	delete(config.Config.Hosts, name)
	return nil
}
