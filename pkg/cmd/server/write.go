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
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

var cmdWrite = &cobra.Command{
	Use:   "write",
	Short: "Write to a specified device",
	Long: heredoc.Doc(`
		Write data to a specified device managed by Synse Server.

		Writes are routed from Synse Server to the appropriate managing
		plugin. All writes are asynchronous and will return a transaction
		ID. This ID can be checked to get the state and status of the write
		using the 'synse server transaction' command.

		An ACTION and DATA can be supplied to the write. The action is always
		required; data may be optional, depending on whether the plugin
		needs it or not.

		For more information, see:
		https://vapor-ware.github.io/synse-server/#write
	`),

	Run: func(cmd *cobra.Command, args []string) {
		// todo
		fmt.Println("< server write")
	},
}
