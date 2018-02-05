package hosts

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/config"
)

// hostsAddCommand is the CLI sub-command for adding a new host to the CLI
// configuration.
var hostsAddCommand = cli.Command{
	Name:   "add",
	Usage:  "Add a new Synse Server host",
	Action: cmdAdd,
}

// cmdAdd is the action for hostsAddCommand. It adds the specified host to the
// CLI configuration.
func cmdAdd(c *cli.Context) error {
	name := c.Args().Get(0)
	addr := c.Args().Get(1)
	if name == "" || addr == "" {
		return cli.NewExitError("'add' requires 2 arguments", 1)
	}

	err := config.Config.AddHost(config.NewHostConfig(name, addr))
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	return nil
}
