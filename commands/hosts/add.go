package hosts

import (
	"fmt"

	"github.com/urfave/cli"
)

var hostAddCommand = cli.Command{
	Name: "add",
	Usage: "add a Synse Server instance to the tracked hosts",
	Action: cmfAdd,

}

func cmfAdd(c *cli.Context) error {
	fmt.Println("hosts add")
	return nil
}