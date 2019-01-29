package change

import (
	"fmt"

	"github.com/spf13/cobra"
)

// New returns a new instance of the 'hosts change' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "change",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< hosts change")
		},
	}

	return cmd
}
