// Synse CLI
// Copyright (c) 2019 Vapor IO
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	ctx "context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/cmd/context"
	"github.com/vapor-ware/synse-cli/pkg/cmd/plugin"
	"github.com/vapor-ware/synse-cli/pkg/cmd/server"
	"github.com/vapor-ware/synse-cli/pkg/cmd/template"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-cli/pkg/utils/exit"
	"github.com/vapor-ware/synse-cli/pkg/view"
)

func init() {
	// Set the logging level to panic, which effectively disables logging.
	// It can be enabled with the debug flag.
	log.SetLevel(log.PanicLevel)

	rootCmd.PersistentFlags().BoolVarP(&flagDebug, "debug", "d", false, "enable debug logging")
}

var (
	flagDebug  bool
	flagSimple bool
)

// resetFlags resets the flag values. This is useful for tests.
func resetFlags() {
	flagDebug = false
	flagSimple = false
}

// rootCmd is the root command for synse.
var rootCmd = &cobra.Command{
	Use:   "synse",
	Short: "Command-line interface for the Synse platform",
	Long: utils.Doc(`
		Command-line interface for the Synse platform.

		Synse is a platform for monitoring and controlling physical and virtual
		devices at data center scale.

		This tool provides simple access to Synse APIs as well as simple
		management and development utilities for the Synse platform.

		<underscore>https://github.com/vapor-ware/synse</>
	`),
	BashCompletionFunction: bashCompletionFunc,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if flagDebug {
			log.SetLevel(log.DebugLevel)
		}

		// Load CLI config from file prior to running any command.
		exit.FromCmd(cmd).Err(config.Load())

		log.WithFields(log.Fields{
			"command": cmd.Name(),
			"args":    args,
		}).Debug("running command")
	},
	Run: run,

	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// Persist the CLI config to file after running any command.
		exit.FromCmd(cmd).Err(config.Persist())
	},
}

func init() {
	rootCmd.AddCommand(
		context.New(),
		plugin.New(),
		server.New(),
		template.New(),

		cmdCompletion,
		cmdVersion,
	)
}

// Execute runs the root command; the entry point into the Synse CLI.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.WithField("error", err).Error("error running root command")
	}
}

func run(_ *cobra.Command, _ []string) {
	cx, cancel := ctx.WithCancel(ctx.Background())
	defer cancel()

	c, err := client.NewAPIClient(cx, "", "")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	i := view.NewInstance()
	i.APIClient = c

	if err := i.Init(); err != nil {
		panic(fmt.Sprintf("failed to init synse-cli: %v", err))
	}
	if err := i.Run(); err != nil {
		panic(fmt.Sprintf("failed to start synse-cli: %v", err))
	}
}
