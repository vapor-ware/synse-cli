package scan

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	filters []string
	sort string
)

// New returns a new instance of the 'server scan' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan",
		Short: "Enumerate all devices",
		Long:  `Enumerate all devices available to Synse Server.

Scan results can be sorted and filtered. This is particularly useful
when a Synse Server instance manages many devices.

Sorting is done via the '--sort' flag. The value for the flag specifies
the fields that should be sorted. Multiple fields can be specified in
a comma-separated string, e.g. "rack,board". The first field will is the
primary sort key, the second field is the secondary sort key, etc.
The following fields support sorting:
 - rack   - device
 - board  - type

Filtering is done via the '--filter' flag. The value for the flag is
a string in the format "KEY=VALUE", where KEY is the field to filter
by, and VALUE is the desired value to filter against. The '--filter'
flag can be used multiple times to specify multiple filters. The
following fields support filtering:
 - rack   - type
 - board

Some examples:
* Show only LED devices sorted by their rack, board, and device ids:
  sysne server scan --sort "rack,board,device" --filter "type=led"

* Show only temperature and pressure devices:
  synse server scan --filter "type=temperature" --filter "type=pressure"

For more information, see:
https://vapor-ware.github.io/synse-server/#scan`,

		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< server scan")
		},
	}

	// Add flag options to the command.
	cmd.PersistentFlags().StringArrayVar(&filters, "filter", []string{}, "set filter(s) for the output results")
	cmd.PersistentFlags().StringVar(&sort, "sort", "", "set the sorting constraints")

	return cmd
}
