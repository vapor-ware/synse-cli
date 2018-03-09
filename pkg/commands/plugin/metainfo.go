package plugin

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// metainfoCmdName is the name for the 'metainfo' command.
	metainfoCmdName = "meta"

	// metainfoCmdUsage is the usage text for the 'metainfo' command.
	metainfoCmdUsage = "Get meta information from a plugin"

	// metainfoCmdDesc is the description for the 'metainfo' command.
	metainfoCmdDesc = `The meta command gets meta information from a plugin via the
  Synse gRPC API. The plugin meta info return is similar to that
  of a 'synse server scan' command, but contains more information
  about the device.

  The 'plugin meta' command takes no arguments.

Example:
  synse plugin meta

Formatting:
  The 'plugin meta' command supports the following formatting
  options (via the CLI global --format flag):
    - pretty (default)`
)

// pluginMetainfoCommand is a CLI sub-command for getting metainfo from a plugin.
var pluginMetainfoCommand = cli.Command{
	Name:        metainfoCmdName,
	Usage:       metainfoCmdUsage,
	Description: metainfoCmdDesc,
	ArgsUsage:   utils.NoArgs,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdMeta(c))
	},
}

// cmdMeta is the action for pluginMetainfoCommand. It prints out the meta-information
// provided by the specified plugin.
func cmdMeta(c *cli.Context) error {

	resp, err := client.Grpc.Metainfo(c)
	if err != nil {
		return err
	}

	formatter := formatters.NewMetaFormatter(c)
	for _, meta := range resp {
		err = formatter.Add(meta)
		if err != nil {
			return err
		}
	}
	return formatter.Write()
}
