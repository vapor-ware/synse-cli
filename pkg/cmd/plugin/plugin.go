package plugin

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/cmd/plugin/capabilities"
	"github.com/vapor-ware/synse-cli/pkg/cmd/plugin/devices"
	"github.com/vapor-ware/synse-cli/pkg/cmd/plugin/health"
	"github.com/vapor-ware/synse-cli/pkg/cmd/plugin/read"
	"github.com/vapor-ware/synse-cli/pkg/cmd/plugin/readcached"
	"github.com/vapor-ware/synse-cli/pkg/cmd/plugin/template"
	"github.com/vapor-ware/synse-cli/pkg/cmd/plugin/transaction"
	"github.com/vapor-ware/synse-cli/pkg/cmd/plugin/write"
)

// New returns a new instance of the 'plugin' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plugin",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< plugin")
		},
	}

	// Add sub-commands
	cmd.AddCommand(
		capabilities.New(),
		devices.New(),
		health.New(),
		read.New(),
		readcached.New(),
		template.New(),
		transaction.New(),
		write.New(),
	)

	return cmd
}
