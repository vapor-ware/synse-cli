package completion

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/config"
)

// CompleteHostNames is a bash completion function for hosts commands. It will
// auto-complete on the names of the configured Synse Server hosts, if any exist.
func CompleteHostNames(c *cli.Context) {
	if c.NArg() > 0 {
		return
	}
	for name := range config.Config.Hosts {
		fmt.Println(name)
	}
}

// CompleteRackBoardDevicePlugin is a bash completion function for plugin commands. It
// will auto-complete on the rack, board, and device ids returned from a plugin
// "metainfo" request.
func CompleteRackBoardDevicePlugin(c *cli.Context) { // nolint: gocyclo
	meta, err := client.Grpc.Metainfo(c)
	if err != nil {
		return
	}

	// If there are no arguments, resolve the first arg, rack
	if c.NArg() == 0 {
		// To get the unique rack id strings, we create a map where the keys will
		// be the rack id and the value just an empty struct. The value is ignored.
		set := make(map[string]struct{})
		for _, device := range meta {
			set[device.Location.Rack] = struct{}{}
		}

		for opt := range set {
			fmt.Println(opt)
		}
	}

	// If there is one argument, resolve the second arg, board
	if c.NArg() == 1 {
		// Get the rack ID being used
		rackID := c.Args().Get(0)

		// To get the unique board id strings, we create a map where the keys will
		// be the board id and the value just an empty struct. The value is ignored.
		set := make(map[string]struct{})
		for _, device := range meta {
			if device.Location.Rack == rackID {
				set[device.Location.Board] = struct{}{}
			}
		}

		for opt := range set {
			fmt.Println(opt)
		}
	}

	// If there are two arguments, resolve the third arg, device
	if c.NArg() == 2 {
		// Get the rack ID and board ID being used
		rackID := c.Args().Get(0)
		boardID := c.Args().Get(1)

		// To get the unique device id strings, we create a map where the keys will
		// be the device id and the value just an empty struct. The value is ignored.
		set := make(map[string]struct{})
		for _, device := range meta {
			if device.Location.Rack == rackID && device.Location.Board == boardID {
				set[device.Uid] = struct{}{}
			}
		}

		for opt := range set {
			fmt.Println(opt)
		}
	}

	// If there are three or more args, we are done, so return.
	if c.NArg() >= 3 {
		return
	}
}

// CompleteRackBoardDevice is a bash completion function for server commands. It will
// auto-complete on the rack, board, and device ids returned from a "scan" request.
func CompleteRackBoardDevice(c *cli.Context) { // nolint: gocyclo
	scan, err := client.Client.Scan()
	if err != nil {
		return
	}
	// Convert to an internal representation to make it easier to
	// do resource (rack, board, device) aggregation.
	devices := scan.ToScanDevices()

	// If there are no arguments, resolve the first arg, rack
	if c.NArg() == 0 {
		// To get the unique rack id strings, we create a map where the keys will
		// be the rack id and the value just an empty struct. The value is ignored.
		set := make(map[string]struct{})
		for _, device := range devices {
			set[device.Rack] = struct{}{}
		}

		for opt := range set {
			fmt.Println(opt)
		}
	}

	// If there is one argument, resolve the second arg, board
	if c.NArg() == 1 {
		// Get the rack ID being used
		rackID := c.Args().Get(0)

		// To get the unique board id strings, we create a map where the keys will
		// be the board id and the value just an empty struct. The value is ignored.
		set := make(map[string]struct{})
		for _, device := range devices {
			if device.Rack == rackID {
				set[device.Board] = struct{}{}
			}
		}

		for opt := range set {
			fmt.Println(opt)
		}
	}

	// If there are two arguments, resolve the third arg, device
	if c.NArg() == 2 {
		// Get the rack ID and board ID being used
		rackID := c.Args().Get(0)
		boardID := c.Args().Get(1)

		// To get the unique device id strings, we create a map where the keys will
		// be the device id and the value just an empty struct. The value is ignored.
		set := make(map[string]struct{})
		for _, device := range devices {
			if device.Rack == rackID && device.Board == boardID {
				set[device.Device] = struct{}{}
			}
		}

		for opt := range set {
			fmt.Println(opt)
		}
	}

	// If there are three or more args, we are done, so return.
	if c.NArg() >= 3 {
		return
	}
}

// CompleteTransactions is a bash completion function for the server 'transaction' command.
// It will auto-complete on the transaction ids returned from a "transaction" request.
func CompleteTransactions(c *cli.Context) {
	// If there are no arguments, resolve the transaction ID
	if c.NArg() == 0 {
		ids, err := client.Client.TransactionList()
		if err != nil {
			return
		}

		for _, opt := range *ids {
			fmt.Println(opt)
		}
	}
}
