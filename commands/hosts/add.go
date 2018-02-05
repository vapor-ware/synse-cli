package hosts

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/config"
)

var hostAddCommand = cli.Command{
	Name: "add",
	Usage: "add a Synse Server instance to the tracked hosts",
	Action: cmfAdd,

}

func cmfAdd(c *cli.Context) error {
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