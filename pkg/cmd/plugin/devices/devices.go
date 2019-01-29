package devices

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'plugin devices' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "devices",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< plugin devices")
		},
	}

	return cmd
}
