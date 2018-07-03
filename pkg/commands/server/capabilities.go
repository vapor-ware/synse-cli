package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// capabilitiesCmdName is the name for the 'capabilities' command.
	capabilitiesCmdName = "capabilities"

	// capabilitiesCmdUsage is the usage text for the 'capabilities' command.
	capabilitiesCmdUsage = "Enumerate all devices' capabilities of configured plugins on the active host"

	// capabilitiesCmdDesc is the description for the 'capabilities' command.
	capabilitiesCmdDesc = `The capabilities command hits the active Synse Server host's '/capabilities'
  endpoint, which enumerates all devices' capabilities of configured plugins that are known to Synse
  Server.

Example:
  synse server capabilities 

Formatting:
  The 'server capabilities' command supports the following formatting
  options (via the CLI global --format flag):
    - pretty (default)
    - yaml
    - json`
)

// capabilitiesCommand is the CLI command for Synse Server's "capabilities" API route.
var capabilitiesCommand = cli.Command{
	Name:        capabilitiesCmdName,
	Usage:       capabilitiesCmdUsage,
	Description: capabilitiesCmdDesc,
	ArgsUsage:   utils.NoArgs,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdCapabilities(c))
	},
}

// cmdCapabilities is the action for the capabilitiesCommand. It makes a
// "capabilities" request against the active Synse Server instance.
func cmdCapabilities(c *cli.Context) error {
	capabilities, err := client.Client.Capabilities()
	if err != nil {
		return err
	}

	formatter := formatters.NewServerCapabilitiesFormatter(c)
	err = formatter.Add(capabilities)
	if err != nil {
		return err
	}
	return formatter.Write()
}
