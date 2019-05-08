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
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

// New returns a new instance of the 'plugin' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plugin",
		Short: "Issue commands to Synse plugins",
		Long: heredoc.Doc(`
			Issue commands to Synse plugins via the Synse gRPC API.
		`),
		Run: func(cmd *cobra.Command, args []string) {
			// todo
			fmt.Println("< plugin")
		},
	}

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
