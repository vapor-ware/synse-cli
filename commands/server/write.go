package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/formatters"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
)

const (
	// writeBase is the base URI for the 'write' route.
	writeBase = "write"

	// writeCmdName is the name for the 'write' command.
	writeCmdName = "write"

	// writeCmdUsage is the usage text for the 'write' command.
	writeCmdUsage = "Write to the specified device"

	// writeCmdDesc is the description for the 'write' command.
	writeCmdDesc = `The write command hits the active Synse Server host's '/write'
	 endpoint to write to the specified device. A Synse Server write
	 will be passed along to the backend plugin which handles the
	 given device to get the write information. Not all devices
	 may support writing; device support for write is specified at
	 the plugin level.`
)

// writePost defines the data to POST to the Synse Server "write" route.
type writePost struct {
	Action string `json:"action,omitempty"`
	Raw    string `json:"raw,omitempty"`
}

// WriteCommand is the CLI command for Synse Server's "write" API route.
var WriteCommand = cli.Command{
	Name:        writeCmdName,
	Usage:       writeCmdUsage,
	Description: writeCmdDesc,
	Category:    SynseActionsCategory,
	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdWrite(c))
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
	uri := client.MakeURI(writeBase, rack, board, device)
	err = client.DoPost(uri, body, &write)
	if err != nil {
		return err
	}

	formatter := formatters.NewWriteFormatter(c.App.Writer)
	err = formatter.Add(write)
	if err != nil {
		return err
	}
	return formatter.Write()
}
