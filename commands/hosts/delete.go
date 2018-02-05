package hosts

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/config"
)

var hostDeleteCommand = cli.Command{
	Name:   "delete",
	Usage:  "delete a Synse Server instance configuration",
	Action: cmdDelete,
}

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
