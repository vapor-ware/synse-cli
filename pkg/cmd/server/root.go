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

package server

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/cmd/server/plugins"
)

// New returns a new instance of the 'server' command.
func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Issue commands to Synse Server",
		Long:  heredoc.Doc(`Issue commands to Synse Server via its HTTP API.`),
	}

	// Add sub-commands
	cmd.AddCommand(
		plugins.New(),
		cmdConfig,
		cmdInfo,
		cmdRead,
		cmdReadCache,
		cmdScan,
		cmdStatus,
		cmdTags,
		cmdTransaction,
		cmdVersion,
		cmdWrite,
	)

	return cmd
}
