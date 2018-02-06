package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/flags"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
)

// testBase is the base URI for the "test" route.
const testBase = "test"

// StatusCommand is the CLI command for Synse Server's "test" API route.
var StatusCommand = cli.Command{
	Name:     "status",
	Usage:    "Get the status of the active host",
	Category: "Synse Server Actions",
	Flags: []cli.Flag{
		flags.OutputFlag,
	},
	Action: func(c *cli.Context) error {
		return utils.CmdHandler(c, cmdStatus(c))
	},
}

// cmdStatus is the action for the StatusCommand. It makes an "status" request
// against the active Synse Server instance.
func cmdStatus(c *cli.Context) error {
	status := &scheme.TestStatus{}
	err := utils.DoGetUnversioned(testBase, status)
	if err != nil {
		return err
	}

	return utils.FormatOutput(c, status)
}
