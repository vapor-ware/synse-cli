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

package context

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

var cmdEdit = &cobra.Command{
	Use:   "edit",
	Short: "Edit a context record",
	Long: heredoc.Doc(`
		Edit the context configuration records.

		This allows you to update information about a context, such as the name
		or address.
	`),
	Run: func(cmd *cobra.Command, args []string) {
		// todo: implement editing.
		// see: https://github.com/kubernetes/kubernetes/blob/aef117999658b24628c6ba49685b4d5ca5998308/pkg/kubectl/cmd/util/editor/editor.go
		// for reference around how kubectl does this.
		fmt.Println("not yet implemented")
	},
}
