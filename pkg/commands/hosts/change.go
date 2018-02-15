package hosts

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

// hostsChangeCommand is the CLI sub-command for changing the active host.
var hostsChangeCommand = cli.Command{
	Name:  "change",
	Usage: "Change the active host",

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdChange(c))
	},

	BashComplete: cmdChangeComplete,
}

// cmdChange is the action for hostsChangeCommand. It changes the active host to
// the specified host, if it exists.
func cmdChange(c *cli.Context) error {
	err := utils.RequiresArgsExact(1, c)
	if err != nil {
		return err
	}

	name := c.Args().Get(0)

	for _, host := range config.Config.Hosts {
		if host.Name == name {
			config.Config.ActiveHost = host
			return nil
		}
	}
	return cli.NewExitError(fmt.Sprintf("host with name '%v' not found", name), 1)
}

// cmdChangeComplete is the bash completion function for the hosts change command.
// It will auto-complete on the names of the configured Synse Server hosts, if any
// exist.
func cmdChangeComplete(c *cli.Context) {
	if c.NArg() > 0 {
		return
	}
	for name, _ := range config.Config.Hosts {
		fmt.Println(name)
	}
}
