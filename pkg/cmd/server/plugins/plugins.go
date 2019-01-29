package plugins

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/cmd/server/plugins/health"
	"github.com/vapor-ware/synse-cli/pkg/cmd/server/plugins/info"
)

// New returns a new instance of the 'server plugin' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plugins",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< server plugins")
		},
	}

	// Add sub-commands
	cmd.AddCommand(
		health.New(),
		info.New(),
	)

	return cmd
}
