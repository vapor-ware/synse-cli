package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'server version' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Get the version of Synse Server",
		Long:  `Get the version information for the configured Synse Server instance.

Both the full semantic version and the API version are returned.

For more information, see:
https://vapor-ware.github.io/synse-server/#version`,

		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< server version")
		},
	}

	return cmd
}
