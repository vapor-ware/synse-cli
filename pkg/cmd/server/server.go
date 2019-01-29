package server

import (
	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/cmd/server/capabilities"
	"github.com/vapor-ware/synse-cli/pkg/cmd/server/config"
	"github.com/vapor-ware/synse-cli/pkg/cmd/server/plugins"
	"github.com/vapor-ware/synse-cli/pkg/cmd/server/plugins/info"
	"github.com/vapor-ware/synse-cli/pkg/cmd/server/read"
	"github.com/vapor-ware/synse-cli/pkg/cmd/server/readcached"
	"github.com/vapor-ware/synse-cli/pkg/cmd/server/scan"
	"github.com/vapor-ware/synse-cli/pkg/cmd/server/status"
	"github.com/vapor-ware/synse-cli/pkg/cmd/server/transaction"
	"github.com/vapor-ware/synse-cli/pkg/cmd/server/version"
	"github.com/vapor-ware/synse-cli/pkg/cmd/server/write"
)

// New returns a new instance of the 'server' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "",
		Long:  ``,
	}

	// Add sub-commands
	cmd.AddCommand(
		capabilities.New(),
		config.New(),
		info.New(),
		plugins.New(),
		read.New(),
		readcached.New(),
		scan.New(),
		status.New(),
		transaction.New(),
		version.New(),
		write.New(),
	)

	return cmd
}
