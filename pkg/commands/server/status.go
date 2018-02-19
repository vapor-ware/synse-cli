package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// testCmdName is the name for the 'status' command.
	testCmdName = "status"

	// testCmdUsage is the usage text for the 'status' command.
	testCmdUsage = "Get the status of the active host"

	// testCmdDesc is the description for the 'status' command.
	testCmdDesc = `The status command hits the active Synse Server host's '/test'
	 endpoint, which returns the status of the instance. If the returned
	 status is "ok", then Synse Server is up and reachable. Otherwise there
	 is an error either with Synse Server or connecting to it.`
)

// StatusCommand is the CLI command for Synse Server's "test" API route.
var StatusCommand = cli.Command{
	Name:        testCmdName,
	Usage:       testCmdUsage,
	Description: testCmdDesc,
	Category:    SynseActionsCategory,
	ArgsUsage:   utils.NoArgs,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdStatus(c))
	},
}

// cmdStatus is the action for the StatusCommand. It makes an "status" request
// against the active Synse Server instance.
func cmdStatus(c *cli.Context) error {
	status, err := client.Client.Status()
	if err != nil {
		return err
	}

	formatter := formatters.NewStatusFormatter(c, status)
	return formatter.Write()
}
