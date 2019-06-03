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

package plugin

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
	flagIds      bool
	flagWait     bool
	flagNS       string
	flagStart    string
	flagEnd      string
	flagTags     []string

	flagTLSCert string
	flagContext string
)

var exitutil utils.Exiter

func init() {
	exitutil = &utils.DefaultExiter{}
}

// New returns a new instance of the 'plugin' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plugin",
		Short: "Issue commands to Synse plugins",
		Long: utils.Doc(`
			Issue commands to Synse plugins via the Synse gRPC API.

			In order to issue commands to a plugin instance, there must be
			a current plugin context. See 'synse context' for details.
		`),
	}

	// Add flag options
	cmd.PersistentFlags().StringVarP(&flagTLSCert, "tlscert", "", "", "path to TLS certificate file (e.g. ./plugin.pem)")
	cmd.PersistentFlags().StringVarP(&flagTLSCert, "with-context", "", "", "the name of the plugin context to use")

	// Add sub-commands
	cmd.AddCommand(
		cmdDevices,
		cmdHealth,
		cmdMetadata,
		cmdRead,
		cmdReadCache,
		cmdTest,
		cmdTransaction,
		cmdVersion,
		cmdWrite,
	)

	return cmd
}
