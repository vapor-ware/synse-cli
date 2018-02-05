package hosts

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/config"
)

var hostChangeCommand = cli.Command{
	Name:   "change",
	Usage:  "change the Synse Server instance to interface with",
	Action: cmdChange,
}

func cmdChange(c *cli.Context) error {
	name := c.Args().Get(0)
	if name == "" {
		return cli.NewExitError("'change' requires 1 argument", 1)
	}

	for _, host := range config.Config.Hosts {
		if host.Name == name {
			config.Config.ActiveHost = host
			return nil
		}
	}
	return cli.NewExitError(fmt.Sprintf("host with name '%v' not found", name), 1)
}
