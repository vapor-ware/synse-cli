package hosts

import (
	"fmt"

	"github.com/urfave/cli"
)

var hostDeleteCommand = cli.Command{
	Name: "delete",
	Usage: "delete a Synse Server instance configuration",
	Action: cmdDelete,

}

func cmdDelete(c *cli.Context) error {
	fmt.Println("hosts delete")
	return nil
}
