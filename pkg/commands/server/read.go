package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/completion"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// readCmdName is the name for the 'read' command.
	readCmdName = "read"

	// readCmdUsage is the usage text for the 'read' command.
	readCmdUsage = "Read from the specified device"

	// readCmdArgsUsage is the argument usage for the 'read' command.
	readCmdArgsUsage = "[RACK [BOARD [DEVICE]]]"

	// readCmdDesc is the description for the 'read' command.
	readCmdDesc = `The read command hits the active Synse Server host's '/read'
  endpoint to read from the specified device(s). For each device,
  a Synse Server read will be passed along to the backend plugin
  which handles the given device to get the reading information.
  Not all devices may support reading; device support for read is
  specified at the plugin level.

  The 'read' command does not require any further routing information
  to be specified. If no routing info is specified, the CLI will
  read from all devices. This can be a lot of devices, so it is
  recommended to scope the read by providing some level of context.

  If only rack info is provided, the CLI will read all devices on
  that rack. If rack and board are provided, the CLI will read all
  devices on that board. If rack board and device are provided, the
  CLI will read the specified device. Rack, board, and device can
  be found either via 'scan' info or by using tab completion, if
  enabled.

  Filtering by type is done via the '--type' flag. Specifying
  type(s) will filter the devices to be read to match those
  types that are specified. It can be applied to a read at any
  level of context (rack-level, board-level, or device-level).
  If specified at the device level, the device should match the
  specified type(s) or else it will not be read. This flag can
  be specified multiple times to apply multiple type filters,
  for example:
    --type temperature --type pressure

Example:
  # read a specific device
  synse read rack-1 board 29d1a03e8cddfbf1cf68e14e60e5f5cc

  # read all temperature devices on the rack
  synse read rack-1 --type temperature

Formatting:
  The 'server read' command supports the following formatting
  options (via the CLI global --format flag):
    - pretty (default)
    - yaml
    - json`
)

// readCommand is the CLI command for Synse Server's "read" API route.
var readCommand = cli.Command{
	Name:        readCmdName,
	Usage:       readCmdUsage,
	Description: readCmdDesc,
	ArgsUsage:   readCmdArgsUsage,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdRead(c))
	},

	Flags: []cli.Flag{
		// --type, -t flag specifies the type of devices to read from.
		// This flag is applicable to rack, board, and device scope, but
		// if the specified type does not match the device type at the
		// device scope, nothing will be read.
		cli.StringSliceFlag{
			Name:  "type, t",
			Usage: "specify the type of devices to read",
		},
	},

	BashComplete: completion.CompleteRackBoardDevice,
}

type device struct {
	rack   string
	board  string
	device string
}

// cmdRead is the action for the readCommand. It makes a "read" request
// against the active Synse Server instance.
func cmdRead(c *cli.Context) error {
	err := utils.RequiresArgsInRange(0, 3, c)
	if err != nil {
		return err
	}

	rackID := c.Args().Get(0)
	boardID := c.Args().Get(1)
	deviceID := c.Args().Get(2)

	err = validateDevices(rackID, boardID, deviceID)
	if err != nil {
		return err
	}

	devices, err := filterDevices(rackID, boardID, deviceID, c)
	if err != nil {
		return err
	}

	formatter := formatters.NewReadFormatter(c)
	for _, device := range devices {
		read, err := client.Client.Read(device.rack, device.board, device.device)
		if err != nil {
			return err
		}

		err = formatter.Add(read)
		if err != nil {
			return err
		}
	}
	return formatter.Write()
}

// validateDevices checks if the given rack, board, devices are valid, in other
// word, available in the system. If not, it returns the corresponding error
// from Synse Server.
func validateDevices(rackID, boardID, deviceID string) (err error) {
	if boardID == "" {
		_, err = client.Client.RackInfo(rackID)
	} else if deviceID == "" {
		_, err = client.Client.BoardInfo(rackID, boardID)
	} else {
		_, err = client.Client.DeviceInfo(rackID, boardID, deviceID)
	}

	return err
}

// filterDevices is a helper function that takes the given rack, board, and device
// (each can be unspecified) and returns the set of devices that match those search
// parameters. It assumes that the rack, board and device ID are valid.
func filterDevices(rackID, boardID, deviceID string, c *cli.Context) ([]*device, error) { // nolint: gocyclo
	var toRead []*device

	scanResults, err := client.Client.Scan()
	if err != nil {
		return nil, err
	}

	types := c.StringSlice("type")

	for _, rack := range scanResults.Racks {
		if rack.ID == rackID || rackID == "" {
			for _, board := range rack.Boards {
				if board.ID == boardID || boardID == "" {
					for _, dev := range board.Devices {
						if (deviceID != "" && dev.ID == deviceID) || deviceID == "" {
							ok := true
							if len(types) > 0 {
								hasType := false
								for _, typeFilter := range types {
									if dev.Type == typeFilter {
										hasType = true
										break
									}
								}
								if !hasType {
									ok = false
								}
							}

							if ok {
								toRead = append(toRead, &device{
									rack:   rack.ID,
									board:  board.ID,
									device: dev.ID,
								})
							}
						}
					}
				}
			}
		}
	}

	return toRead, nil
}
