package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/flags"
	"github.com/vapor-ware/synse-cli/formatters"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
)

const (
	// versionBase is the base URI for the 'version' route.
	versionBase = "version"

	// versionCmdName is the name for the 'version' command.
	versionCmdName = "version"

	// versionCmdUsage is the usage text for the 'version' command.
	versionCmdUsage = "Get the version of the active host"

	// versionCmdDesc is the description for the 'version' command.
	versionCmdDesc = `The version command hits the active Synse Server host's '/version'
	 endpoint, which returns the version (full and API) of the Synse
	 Server instance.`
)

// VersionCommand is the CLI command for Synse Server's "version" API route.
var VersionCommand = cli.Command{
	Name:        versionCmdName,
	Usage:       versionCmdUsage,
	Description: versionCmdDesc,
	Category:    SynseActionsCategory,
	ArgsUsage:   utils.NoArgs,
	Flags: []cli.Flag{
		flags.OutputFlag,
	},
	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdVersion(c))
	},
}

// cmdVersion is the action for the VersionCommand. It makes an "version" request
// against the active Synse Server instance.
func cmdVersion(c *cli.Context) error {
	version := &scheme.Version{}
	err := client.DoGetUnversioned(versionBase, version)
	if err != nil {
		return err
	}

	return formatters.FormatOutput(c, version)
}
