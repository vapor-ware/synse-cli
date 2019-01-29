package active

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'hosts active' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "active",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< hosts active")
		},
	}

	return cmd
}
