package health

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'plugin health' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "health",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< plugin health")
		},
	}

	return cmd
}
