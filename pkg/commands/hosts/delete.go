package hosts

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"fmt"
)

// hostsDeleteCommand is the CLI sub-command for deleting a host from the
// CLI configuration.
var hostsDeleteCommand = cli.Command{
	Name:  "delete",
	Usage: "Delete a Synse Server host from configuration",

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdDelete(c))
	},

	BashComplete: cmdDeleteComplete,
}

// cmdDelete is the action for hostsDeleteCommand. It removes the specified host
// from the CLI configuration, if it exists. If the specified host is also the
// active host, it will unset the active host.
func cmdDelete(c *cli.Context) error {
	err := utils.RequiresArgsExact(1, c)
	if err != nil {
		return err
	}

	name := c.Args().Get(0)

	host := config.Config.Hosts[name]
	if host != nil {
		if config.Config.ActiveHost != nil && *host == *config.Config.ActiveHost {
			config.Config.ActiveHost = nil
		}
	}
	delete(config.Config.Hosts, name)
	return nil
}

// cmdDeleteComplete is the bash completion function for the hosts delete command.
// It will auto-complete on the names of the configured Synse Server hosts, if any
// exist.
func cmdDeleteComplete(c *cli.Context) {
	if c.NArg() > 0 {
		return
	}
	for name, _ := range config.Config.Hosts {
		fmt.Println(name)
	}
}