package server

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
)

// writeBase is the base URI for the "write" route.
const writeBase = "write"

// writePost defines the data to POST to the Synse Server "write" route.
type writePost struct {
	Action string `json:"action,omitempty"`
	Raw    string `json:"raw,omitempty"`
}

// WriteCommand is the CLI command for Synse Server's "write" API route.
var WriteCommand = cli.Command{
	Name:     "write",
	Usage:    "Write to the specified device",
	Category: "Synse Server Actions",
	Action: func(c *cli.Context) error {
		return utils.CmdHandler(c, cmdWrite(c))
	},
}

// cmdWrite is the action for the WriteCommand. It makes an "write" request
// against the active Synse Server instance.
func cmdWrite(c *cli.Context) error {
	err := utils.RequiresArgsInRange(4, 5, c)
	if err != nil {
		return err
	}

	rack := c.Args().Get(0)
	board := c.Args().Get(1)
	device := c.Args().Get(2)
	action := c.Args().Get(3)
	raw := c.Args().Get(4)

	write := make([]scheme.WriteTransaction, 0)

	body := &writePost{
		Action: action,
		Raw:    raw,
	}
	uri := utils.MakeURI(writeBase, rack, board, device)
	resp, err := client.New().Post(uri).BodyJSON(body).ReceiveSuccess(&write)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return err
	}

	for _, t := range write {
		fmt.Println(t.Transaction)
	}
	return nil
}
