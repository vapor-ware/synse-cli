package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/completion"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// readCachedCmdName is the name for the 'readcached' command.
	readCachedCmdName = "readcached"

	// readCachedCmdUsage is the usage text for the 'readcached' command.
	readCachedCmdUsage = "Stream reading data from all configured plugins"

	// readCachedCmdDesc is the description for the 'readcached' command.
	readCachedCmdDesc = `The readcached command hits the active Synse Server host's
  '/readcached' endpoint to stream reading data from all configured
  plugins.

  The 'readcached' command does not require any further routing information
  to be specified. If no routing info is specified, the CLI will
  stream all the reading data from all available devices. This
  can be a lot of devices, so it is recommended to scope the
  read by providing some level of context, which are timestamp
  in this case.

  Timestamp is formatted in RFC3339/RFC3339Nano and is used
  to specify a bounding on the cache data to return. There
  are two bounding options: starting and ending one, which
  are done via two corresponding flags: '--start' and '--end'.
  If no timestamp is specified, there will not be any bounding
  constraint and the CLI will return all reading data.

Example:
  # stream reading data from all configured plugins
  synse server readcached

  # stream reading data after '2018-11-01T19:13:00.9184028Z'
  synse server readcached --start 2018-11-01T19:13:00.9184028Z


  # stream reading data before '2018-11-01T19:13:00.9184028Z'
  synse server readcached --end 2018-11-11T19:13:00.9184028Z

  # stream reading data within '2018-11-01T19:13:00.9184028Z' and '2018-11-11T19:13:00.9184028Z'
  synse server readcached --start 2018-11-01T19:13:00.9184028Z --end 2018-11-11T19:13:00.9184028Z

Formatting:
  The 'server readcached' command supports the following formatting
  options (via the CLI global --format flag):
    - pretty (default)
    - yaml
    - json`
)

// readCachedCommand is the CLI command for Synse Server's "readcached" API route.
var readCachedCommand = cli.Command{
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

	BashComplete: completion.CompleteRackBoardDevice,
}

// cmdReadCached is the action for the readCachedCommand. It makes a
// "readcached" request against the active Synse Server instance.
func cmdReadCached(c *cli.Context) error {
	params := scheme.ReadCachedParams{
		Start: c.String("start"),
		End:   c.String("end"),
	}

	devices, err := client.Client.ReadCached(params)
	if err != nil {
		return err
	}

	formatter := formatters.NewReadCachedFormatter(c)
	for _, device := range devices {
		err = formatter.Add(device)
		if err != nil {
			return err
		}
	}
	return formatter.Write()
}
