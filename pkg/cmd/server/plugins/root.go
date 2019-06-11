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

package plugins

import (
	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

// Define variables which hold values passed in via flags. These are
// defined here because they are used by multiple commands in the package.
var (
	flagNoHeader bool
	flagJSON     bool
	flagYaml     bool

	flagTLSCert string
	flagContext string
)

// resetFlags resets the flag values. This is useful for tests.
func resetFlags() {
	flagNoHeader = false
	flagJSON = false
	flagYaml = false
	flagTLSCert = ""
	flagContext = ""
}

// New returns a new instance of the 'server plugin' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plugins",
		Short: "Get info about plugins registered with the server",
		Long: utils.Doc(`
			Get information about plugins registered with the Synse Server
			instance.
		`),
	}

	// Add flag options
	cmd.PersistentFlags().StringVarP(&flagTLSCert, "tlscert", "", "", "path to TLS certificate file (e.g. ./server.pem)")
	cmd.PersistentFlags().StringVarP(&flagContext, "with-context", "", "", "the name of the plugin context to use")

	// Add sub-commands
	cmd.AddCommand(
		cmdHealth,
		cmdInfo,
		cmdList,
	)

	return cmd
}
