package write

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'server write' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "write",
		Short: "Write to a specified device",
		Long:  `Write data to a specified device managed by Synse Server.

Writes are routed from Synse Server to the appropriate managing
plugin. All writes are asynchronous and will return a transaction
ID. This ID can be checked to get the state and status of the write
using the 'synse server transaction' command.

An ACTION and DATA can be supplied to the write. The action is always
required; data may be optional, depending on whether the plugin
needs it or not.

For more information, see:
https://vapor-ware.github.io/synse-server/#write`,

		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< server write")
		},
	}

	return cmd
}
