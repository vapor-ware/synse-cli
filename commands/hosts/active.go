package hosts

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/config"
)

var hostsActiveCommand = cli.Command{
	Name: "active",
	Usage: "print the active Synse Server instance",
	Action: cmdActive,

}

func cmdActive(c *cli.Context) error {
	if config.Config.ActiveHost == nil {
		return cli.NewExitError("no active host set", 1)
	}
	fmt.Printf("Name:    %s\n", config.Config.ActiveHost.Name)
	fmt.Printf("Address: %s\n", config.Config.ActiveHost.Address)
	return nil
}
