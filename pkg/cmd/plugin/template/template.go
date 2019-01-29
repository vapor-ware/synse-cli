package template

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'plugin template' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "template",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< plugin template")
		},
	}

	return cmd
}
