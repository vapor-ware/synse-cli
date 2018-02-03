package hosts

import (
	"fmt"

	"github.com/urfave/cli"
)

var hostChangeCommand = cli.Command{
	Name: "change",
	Usage: "change the Synse Server instance to interface with",
	Action: cmdChange,

}

func cmdChange(c *cli.Context) error {
	fmt.Println("hosts change")
	return nil
}
