package plugin

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// readCachedCmdName is the name for the 'readcached' command.
	readCachedCmdName = "readcached"

	// readCachedCmdUsage is the usage text for the 'readcached' command.
	readCachedCmdUsage = "Get cached reading from a plugin"

	// readCachedCmdDesc is the description for the 'readcached' command.
	readCachedCmdDesc = `The readcached command gets a device reading from a plugin via
  the Synse gRPC API. The plugin read info return is similar to that
  of a 'synse server readcached' command, and the response data for
  both should look the same.

  The 'readcached' command does not require any further routing
  information to be specified. If no routing info is specified,
  the CLI will stream all the reading data from all available
  devices. This can be a lot of devices, so it is recommended
  to scope the read by providing some level of context,
  which are timestamp in this case.

  Timestamp is formatted in RFC3339/RFC3339Nano and is used
  to specify a bounding on the cache data to return. There
  are two bounding options: starting and ending one, which
  are done via two corresponding flags: '--start' and '--end'.
  If no timestamp is specified, there will not be any bounding
  constraint and the CLI will return all reading data.

Example:
  # read all cached reading from a specific plugin via tcp
  synse plugin --tcp localhost:50001 readcached

  # read cached reading after '2018-11-01T19:13:00.9184028Z'
  synse plugin --tcp localhost:50001 readcached --start 2018-11-01T19:13:00.9184028Z

Formatting:
  The 'plugin readcached' command supports the following formatting
  options (via the CLI global --format flag):
    - pretty (default)
    - yaml
    - json`
)

// pluginReadCachedCommand is a CLI sub-command for getting cached reading from a plugin.
var pluginReadCachedCommand = cli.Command{
	Name:        readCachedCmdName,
	Usage:       readCachedCmdUsage,
	Description: readCachedCmdDesc,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdReadCached(c))
	},

	Flags: []cli.Flag{
		// --start, -s flag specifies specifies a starting bound on the cache
		// data to return. If the specified timestamp does not exist, nothing
		// will be read.
		cli.StringFlag{
			Name:  "start, s",
			Usage: "specify the starting bound on the cache data",
		},
		// --end, -e flag specifies specifies a ending bound on the cache
		// data to return. If the specified timestamp does not exist, nothing
		// will be read.
		cli.StringFlag{
			Name:  "end, e",
			Usage: "specify the ending bound on the cache data",
		},
	},
}

// cmdReadCached is the action for pluginReadCachedCommand. It prints out
// cached reading that was retrieved from the specified plugin.
func cmdReadCached(c *cli.Context) error { // nolint: gocyclo
	resp, err := client.Grpc.ReadCached(c, c.String("start"), c.String("end"))
	if err != nil {
		return err
	}

	formatter := formatters.NewReadCachedFormatter(c)
	for _, device := range resp {
		err = formatter.Add(scheme.ReadCached{
			Location: scheme.DeviceLocation{
				Rack:   device.GetRack(),
				Board:  device.GetBoard(),
				Device: device.GetDevice(),
			},
			ReadData: scheme.ReadData{
				Info:      device.GetReading().GetInfo(),
				Type:      device.GetReading().GetType(),
				Timestamp: device.GetReading().GetTimestamp(),
				Unit: scheme.OutputUnit{
					Name:   device.GetReading().GetUnit().GetName(),
					Symbol: device.GetReading().GetUnit().GetSymbol(),
				},
				Value: GetValue(device.GetReading()),
			},
		})
		if err != nil {
			return err
		}
	}
	return formatter.Write()
}
