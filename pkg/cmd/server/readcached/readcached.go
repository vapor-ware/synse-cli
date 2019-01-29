package readcached

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	start string
	end string
)

// New returns a new instance of the 'server readcached' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "readcached",
		Short: "Get cached readings from all configured plugins",
		Long:  `Get the cached readings from all of the plugins registered with Synse Server.

This command operates on all plugins, so it does not require any routing
information to be specified. Start and end timestamps can be set in order
to bound the reading data. It is suggested to use timestamp bounding to
keep output manageable.

The start and end timestamps should be formatted in RFC3339 or RFC3339Nano
format. If no bounding timestamps are specified, all readings will be
returned.

For more information, see:
https://vapor-ware.github.io/synse-server/#read-cached`,

		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< server readcached")
		},
	}

	// Add flag options to the command.
	cmd.PersistentFlags().StringVar(&start, "start", "", "the starting timestamp bound for the read cache data")
	cmd.PersistentFlags().StringVar(&end, "end", "", "the ending timestamp bound for the read cache data")

	return cmd
}
