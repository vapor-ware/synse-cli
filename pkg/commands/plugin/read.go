package plugin

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/completion"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-server-grpc/go"
)

const (
	// readCmdName is the name for the 'read' command.
	readCmdName = "read"

	// readCmdUsage is the usage text for the 'read' command.
	readCmdUsage = "Get a reading from a plugin"

	// readCmdArgsUsage is the argument usage for the 'read' command.
	readCmdArgsUsage = "RACK BOARD DEVICE"

	// readCmdDesc is the description for the 'read' command.
	readCmdDesc = `The read command gets a device reading from a plugin via the
  Synse gRPC API. The plugin read info return is similar to that
  of a 'synse server read' command, and the response data for
  both should look the same.

  Reads require the rack, board, and device to be specified.
  These can be found using the 'plugin meta' command, or with
  tab completion, if enabled.

Example:
  synse plugin read rack-1 board 29d1a03e8cddfbf1cf68e14e60e5f5cc

Formatting:
  The 'plugin read' command supports the following formatting
  options (via the CLI global --format flag):
    - pretty (default)
		- yaml
    - json`
)

// pluginReadCommand is a CLI sub-command for getting a reading from a plugin.
var pluginReadCommand = cli.Command{
	Name:        readCmdName,
	Usage:       readCmdUsage,
	Description: readCmdDesc,
	ArgsUsage:   readCmdArgsUsage,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdRead(c))
	},

	BashComplete: completion.CompleteRackBoardDevicePlugin,
}

// cmdRead is the action for pluginReadCommand. It prints out a reading that was
// retrieved from the specified plugin.
func cmdRead(c *cli.Context) error { // nolint: gocyclo
	err := utils.RequiresArgsExact(3, c)
	if err != nil {
		return err
	}

	rack := c.Args().Get(0)
	board := c.Args().Get(1)
	device := c.Args().Get(2)

	resp, err := client.Grpc.Read(c, rack, board, device)
	if err != nil {
		return err
	}

	readData := make([]scheme.ReadData, 0)
	for _, read := range resp {
		readData = append(readData, scheme.ReadData{
			Timestamp: read.GetTimestamp(),
			Type:      read.GetType(),
			Info:      read.GetInfo(),
			Value:     getValue(read),
			Unit:      scheme.OutputUnit{read.Unit.Name, read.Unit.Symbol},
		})
	}

	formatter := formatters.NewReadFormatter(c)
	err = formatter.Add(&scheme.Read{Data: readData})
	if err != nil {
		return err
	}

	return formatter.Write()
}

// getValue determines which oneof Reading.Value is set and returns the
// corresponding value.
func getValue(value *synse.Reading) interface{} { // nolint: gocyclo
	switch value.Value.(type) {
	case *synse.Reading_StringValue:
		return value.GetStringValue()
	case *synse.Reading_BoolValue:
		return value.GetBoolValue()
	case *synse.Reading_Float32Value:
		return value.GetFloat32Value()
	case *synse.Reading_Float64Value:
		return value.GetFloat64Value()
	case *synse.Reading_Int32Value:
		return value.GetInt32Value()
	case *synse.Reading_Int64Value:
		return value.GetInt64Value()
	case *synse.Reading_BytesValue:
		return value.GetBytesValue()
	case *synse.Reading_Uint32Value:
		return value.GetUint32Value()
	case *synse.Reading_Uint64Value:
		return value.GetUint64Value()
	default:
		// FIXME: Should we return nil here?
		return nil
	}
}
