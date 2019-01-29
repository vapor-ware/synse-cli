package delete

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'hosts delete' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< hosts delete")
		},
	}

	return cmd
}
