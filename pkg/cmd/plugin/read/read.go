package read

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'plugin read' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "read",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< plugin read")
		},
	}

	return cmd
}
