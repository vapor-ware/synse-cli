package hosts

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/config"
)

// hostsActiveCommand is the CLI sub-command for getting the active host.
var hostsActiveCommand = cli.Command{
	Name:   "active",
	Usage:  "Display information for the active host",
	Action: cmdActive,
}

// cmdActive is the action for hostsActiveCommand. It prints out the information
// associated with the active host, if one is set.
func cmdActive(c *cli.Context) error {
	if config.Config.ActiveHost == nil {
		return cli.NewExitError("no active host set", 1)
	}
	fmt.Printf("Name:    %s\n", config.Config.ActiveHost.Name)
	fmt.Printf("Address: %s\n", config.Config.ActiveHost.Address)
	return nil
}
