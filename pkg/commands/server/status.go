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
	testCmdUsage = "Get the status of Synse Server"

	// testCmdDesc is the description for the 'status' command.
	testCmdDesc = `The status command hits the active Synse Server host's '/test'
  endpoint, which returns the status of the instance. If the returned
  status is "ok", then Synse Server is up and reachable. Otherwise there
  is an error either with Synse Server or connecting to it.

  The 'synse status' command takes no arguments.

Example:
  synse server status

Formatting:
  The 'server status' command supports the following formatting
  options (via the CLI global --format flag):
    - yaml (default)
    - json`
)

// statusCommand is the CLI command for Synse Server's "test" API route.
var statusCommand = cli.Command{
	Name:        testCmdName,
	Usage:       testCmdUsage,
	Description: testCmdDesc,
	ArgsUsage:   utils.NoArgs,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdStatus(c))
	},
}

// cmdStatus is the action for the statusCommand. It makes a "status" request
// against the active Synse Server instance.
func cmdStatus(c *cli.Context) error {
	status, err := client.Client.Status()
	if err != nil {
		return err
	}

	formatter := formatters.NewStatusFormatter(c)
	err = formatter.Add(status)
	if err != nil {
		return err
	}
	return formatter.Write()
}
