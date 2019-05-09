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

	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

var cmdEdit = &cobra.Command{

	// FIXME: this needs to be implemented. Until then, make this hidden.
	Hidden: true,

	Use:   "edit",
	Short: "Edit a context record",
	Long: utils.Doc(`
		Edit context records in the synse configuration.
	`),
	SuggestFor: []string{
		"update",
	},
	Run: func(cmd *cobra.Command, args []string) {
		// todo: implement editing.
		// see: https://github.com/kubernetes/kubernetes/blob/aef117999658b24628c6ba49685b4d5ca5998308/pkg/kubectl/cmd/util/editor/editor.go
		// for reference around how kubectl does this.
		fmt.Println("not yet implemented")
	},
}
