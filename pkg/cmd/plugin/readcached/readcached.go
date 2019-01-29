package readcached

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'plugin readcached' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "readcached",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< plugin readcached")
		},
	}

	return cmd
}
