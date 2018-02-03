package hosts

import (
	"fmt"

	"github.com/urfave/cli"
)

var hostsListCommand = cli.Command{
	Name: "list",
	Usage: "list the configured Synse Server hosts",
	Action: cmdList,

}

func cmdList(c *cli.Context) error {
	fmt.Println("hosts list")
	return nil
}
