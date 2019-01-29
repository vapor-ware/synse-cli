package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/cmd/completion"
	"github.com/vapor-ware/synse-cli/pkg/cmd/hosts"
	"github.com/vapor-ware/synse-cli/pkg/cmd/plugin"
	"github.com/vapor-ware/synse-cli/pkg/cmd/server"
	"github.com/vapor-ware/synse-cli/pkg/cmd/update"
	"github.com/vapor-ware/synse-cli/pkg/cmd/version"
)

// rootCmd is the root command for synse.
var rootCmd = &cobra.Command{
	Use:   "synse",
	Short: "Monitor and control physical and virtual devices",
	Long: `A platform to monitor and control physical and virtual devices.

This tool provides access to Synse APIs as well as simple management
and utility operations relating to the Synse platform.

For more information, see: https://github.com/vapor-ware/synse`,
}

func init() {
	rootCmd.AddCommand(
		completion.New(rootCmd),
		hosts.New(),
		plugin.New(),
		server.New(),
		update.New(),
		version.New(),
	)
}

// Execute runs the root command; the entry point into the Synse CLI.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// TODO: error out in a consistent way.
		fmt.Println(err)
		os.Exit(1)
	}
}
