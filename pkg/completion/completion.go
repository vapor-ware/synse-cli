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

// CompleteRackBoardDevice is a bash completion function for server commands. It will
// auto-complete on the rack, board, and device ids returned from a "scan" request.
func CompleteRackBoardDevice(c *cli.Context) { // nolint: gocyclo
	scan, err := client.Client.Scan()
	if err != nil {
		return
	}
	// Convert to an internal representation to make it easier to
	// do resource (rack, board, device) aggregation.
	devices := scan.ToInternalScan()

	// If there are no arguments, resolve the first arg, rack
	if c.NArg() == 0 {
		// to get the unique rack id strings, we create a map where
		// the keys will be the rack id and the value just an empty
		// struct. the value is ignored.
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
		// to get the unique board id strings, we create a map where
		// the keys will be the board id and the value just an empty
		// struct. the value is ignored.
		set := make(map[string]struct{})
		for _, device := range devices {
			set[device.Board] = struct{}{}
		}

		for opt := range set {
			fmt.Println(opt)
		}
	}

	// If there are two arguments, resolve the third arg, device
	if c.NArg() == 2 {
		// to get the unique device id strings, we create a map where
		// the keys will be the device id and the value just an empty
		// struct. the value is ignored.
		set := make(map[string]struct{})
		for _, device := range devices {
			set[device.Device] = struct{}{}
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
