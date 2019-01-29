package write

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'plugin write' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "write",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< plugin write")
		},
	}

	return cmd
}
