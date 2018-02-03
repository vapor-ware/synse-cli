package hosts

import (
	"fmt"

	"github.com/urfave/cli"
)

var hostsActiveCommand = cli.Command{
	Name: "active",
	Usage: "print the active Synse Server instance",
	Action: cmdActive,

}

func cmdActive(c *cli.Context) error {
	fmt.Println("hosts active")
	return nil
}
