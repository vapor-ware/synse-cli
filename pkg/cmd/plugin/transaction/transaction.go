package transaction

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'plugin transaction' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transaction",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< plugin transaction")
		},
	}

	return cmd
}
