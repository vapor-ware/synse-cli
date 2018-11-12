package plugin

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// devicesCmdName is the name for the 'devices' command.
	devicesCmdName = "devices"

	// devicesCmdUsage is the usage text for the 'devices' command.
	devicesCmdUsage = "Get device information from a plugin"

	// devicesCmdDesc is the description for the 'devices' command.
	devicesCmdDesc = `The devices command gets device information from a plugin via the
  Synse gRPC API. The device information returned here is similar to that
  of a 'synse server scan' command, but contains more information
  about the device.

  The 'plugin devices' command takes no arguments.

Example:
  synse plugin devices

Formatting:
  The 'plugin devices' command supports the following formatting
  options (via the CLI global --format flag):
    - pretty (default)
    - yaml
    - json`
)

// pluginDevicesCommand is a CLI sub-command for getting devices info from a plugin.
var pluginDevicesCommand = cli.Command{
	Name:        devicesCmdName,
	Usage:       devicesCmdUsage,
	Description: devicesCmdDesc,
	ArgsUsage:   utils.NoArgs,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdDevices(c))
	},
}

// cmdDevices is the action for pluginDevicesCommand. It prints out the devices information
// provided by the specified plugin.
func cmdDevices(c *cli.Context) error {
	resp, err := client.Grpc.Devices(c)
	if err != nil {
		return err
	}

	formatter := formatters.NewDevicesFormatter(c)
	for _, device := range resp {
		err = formatter.Add(device)
		if err != nil {
			return err
		}
	}
	return formatter.Write()
}
