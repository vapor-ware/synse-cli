package server

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
)

// readURI
const readURI = "read"

// readCommand
var readCommand = cli.Command{
	Name:  "read",
	Usage: "read",
	Action: func(c *cli.Context) error {
		return utils.CommandHandler(c, cmdRead(c))
	},
}

// cmdRead
func cmdRead(c *cli.Context) error {
	read := &scheme.Read{}
	resp, err := client.New().Get(readURI).ReceiveSuccess(read)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return err
	}

	fmt.Println("unimplemented")
	return nil
}
