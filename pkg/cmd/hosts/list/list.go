package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'hosts list' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< hosts list")
		},
	}

	return cmd
}
