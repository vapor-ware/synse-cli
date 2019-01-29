package read

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'server read' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "read",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< server read")
		},
	}

	return cmd
}
