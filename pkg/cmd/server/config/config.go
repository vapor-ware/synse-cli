package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'server config' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< server config")
		},
	}

	return cmd
}
