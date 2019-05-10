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
	"os"

	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

var cmdCompletion = &cobra.Command{
	Use:   "completion",
	Short: "Generate completion scripts (bash)",
	Long: utils.Doc(`
		Generate completion scripts (bash).

		To load bash completion for the current session, run:

		  . <(synse completion)

		To configure your bash shell to load synse completion for all
		new sessions, add the above to your bashrc, e.g.

		  echo ". <(synse completion)" >> ~/.bashrc
	`),

	Run: func(cmd *cobra.Command, args []string) {
		exitutil.Err(rootCmd.GenBashCompletion(os.Stdout))
	},
}
