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

package template

import (
	"fmt"

	"github.com/vapor-ware/synse-cli/pkg/utils"

	"github.com/spf13/cobra"
)

var cmdPlugin = &cobra.Command{
	Use:   "plugin",
	Short: "",
	Long: utils.Doc(`

	`),
	Run: func(cmd *cobra.Command, args []string) {
		// todo
		fmt.Println("< template plugin")
	},
}
