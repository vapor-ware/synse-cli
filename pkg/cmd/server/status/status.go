package status

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'server status' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get the status of Synse Server",
		Long:  `Get the connectivity status of a Synse Server instance.

This uses the '/test' endpoint, which is dependency-free and is used
to determine whether the server instance is reachable or not. It does
not provide any other indication of health.

For more information on the status endpoint, see:
https://vapor-ware.github.io/synse-server/#test`,

		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< server status")
		},
	}

	return cmd
}
