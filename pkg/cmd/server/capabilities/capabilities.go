package capabilities

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'server capabilities' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "capabilities",
		Short: "Enumerate device capabilities from registered plugins",
		Long:  `Enumerate the capabilities of devices from each of the registered plugins.

For each device that a plugin supports, it defines the values which can
be read from that device. This command enumerates both the device types
which all registered plugins support as well as the kinds of readings
that each of those devices support.

For more information, see:
https://vapor-ware.github.io/synse-server/#capabilities`,

		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< server capabilities")
		},
	}

	return cmd
}
