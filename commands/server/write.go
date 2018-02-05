package server

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
)

// writeURI
const writeURI = "write"

// writeCommand
var writeCommand = cli.Command{
	Name:  "write",
	Usage: "write",
	Action: func(c *cli.Context) error {
		return utils.CommandHandler(c, cmdWrite(c))
	},
}

// cmdWrite
func cmdWrite(c *cli.Context) error {
	write := &scheme.Write{}
	resp, err := client.New().Get(writeURI).ReceiveSuccess(write)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return err
	}

	fmt.Println("unimplemented")
	return nil
}
