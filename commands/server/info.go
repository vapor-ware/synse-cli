package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/utils"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/client"
	"net/http"
	"fmt"
)

// infoURI
const infoURI = "info"

// infoCommand
var infoCommand = cli.Command{
	Name:    "info",
	Usage:   "info",
	Action: func(c *cli.Context) error {
		return utils.CommandHandler(c, cmdInfo(c))
	},
}

// cmdInfo
func cmdInfo(c *cli.Context) error {
	// FIXME - need to discern between rack, board, and device info
	info := &scheme.RackInfo{}
	resp, err := client.New().Get(infoURI).ReceiveSuccess(info)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return err
	}

	fmt.Println("unimplemented")
	return nil
}
