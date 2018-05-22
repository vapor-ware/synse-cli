package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// versionCmdName is the name for the 'version' command.
	versionCmdName = "version"

	// versionCmdUsage is the usage text for the 'version' command.
	versionCmdUsage = "Get the version of Synse Server"

	// versionCmdDesc is the description for the 'version' command.
	versionCmdDesc = `The version command hits the active Synse Server host's '/version'
  endpoint, which returns the version (full and API) of the Synse
  Server instance.

Example:
  synse server version

Formatting:
  The 'server version' command supports the following formatting
  options (via the CLI global --format flag):
    - yaml (default)
    - json`
)

// versionCommand is the CLI command for Synse Server's "version" API route.
var versionCommand = cli.Command{
	Name:        versionCmdName,
	Usage:       versionCmdUsage,
	Description: versionCmdDesc,
	ArgsUsage:   utils.NoArgs,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdVersion(c))
	},
}

// cmdVersion is the action for the versionCommand. It makes a "version" request
// against the active Synse Server instance.
func cmdVersion(c *cli.Context) error {
	version, err := client.Client.Version()
	if err != nil {
		return err
	}

	formatter := formatters.NewVersionFormatter(c)
	err = formatter.Add(version)
	if err != nil {
		return err
	}
	return formatter.Write()
}
