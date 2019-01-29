package add

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'hosts add' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< hosts add")
		},
	}

	return cmd
}
