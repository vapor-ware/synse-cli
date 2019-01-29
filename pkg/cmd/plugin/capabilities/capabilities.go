package capabilities

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'plugin capabilities' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "capabilities",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< plugin capabilities")
		},
	}

	return cmd
}
