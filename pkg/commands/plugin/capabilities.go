package plugin

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
	capabilitiesCmdUsage = "Get the device capabilities for a plugin"

	// capabilitiesCmdDesc is the description for the 'capabilities' command.
	capabilitiesCmdDesc = `The capabilities command gets information on the kinds of devices
  that a plugin supports as well as the types of reading outputs 
  each of those devices support. The device information returned
  here is similar to that of a 'synse server capabilities' command.

  The 'plugin capabilities' command takes no arguments.

Example:
  synse plugin capabilities

Formatting:
  The 'plugin capabilities' command supports the following formatting
  options (via the CLI global --format flag):
    - pretty (default)`
)

// pluginCapabilitiesCommand is a CLI sub-command for getting capabilities info from a plugin.
var pluginCapabilitiesCommand = cli.Command{
	Name:        capabilitiesCmdName,
	Usage:       capabilitiesCmdUsage,
	Description: capabilitiesCmdDesc,
	ArgsUsage:   utils.NoArgs,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdCapabilities(c))
	},
}

// cmdCapabilities is the action for pluginCapabilitiesCommand. It prints out
// the capabilities information provided by the specified plugin.
func cmdCapabilities(c *cli.Context) error {
	capabilities, err := client.Grpc.Capabilities(c)
	if err != nil {
		return err
	}

	formatter := formatters.NewPluginCapabilitiesFormatter(c)
	for _, capability := range capabilities {
		err = formatter.Add(capability)
		if err != nil {
			return err
		}
	}
	return formatter.Write()
}
