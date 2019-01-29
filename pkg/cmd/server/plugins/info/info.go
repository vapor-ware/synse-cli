package info

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'server plugins info' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< server plugins info")
		},
	}

	return cmd
}
