package server

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
)

// readBase is the base URI for the "read" route.
const readBase = "read"

// ReadCommand is the CLI command for Synse Server's "read" API route.
var ReadCommand = cli.Command{
	Name:     "read",
	Usage:    "Read from the specified device",
	Category: "Synse Server Actions",
	Action: func(c *cli.Context) error {
		return utils.CmdHandler(c, cmdRead(c))
	},
}

// cmdRead is the action for the ReadCommand. It makes an "read" request
// against the active Synse Server instance.
func cmdRead(c *cli.Context) error {
	err := utils.RequiresArgsExact(3, c)
	if err != nil {
		return err
	}

	rack := c.Args().Get(0)
	board := c.Args().Get(1)
	device := c.Args().Get(2)

	read := &scheme.Read{}
	err = utils.DoGet(utils.MakeURI(readBase, rack, board, device), read)
	if err != nil {
		return err
	}

	var data [][]string
	for readType, readData := range read.Data {
		data = append(data, []string{
			readType,
			fmt.Sprintf("%v", readData.Value),
		})
	}

	header := []string{"reading", "value"}
	utils.TableOutput(header, data)
	return nil
}
