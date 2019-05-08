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
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

func init() {
	cmdList.Flags().BoolVarP(&flagFull, "full", "f", false, "display the full context record")
	cmdList.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
}

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "List all configured contexts",
	Long: heredoc.Doc(`
		List all configured contexts.

		This will display all information for each configured context record.
	`),

	Run: func(cmd *cobra.Command, args []string) {
		utils.Err(listContexts())
	},
}

func listContexts() error {
	contexts := config.GetContexts()
	if len(contexts) == 0 {
		return nil
	}

	out := utils.NewTabWriter()
	defer out.Flush()

	if !flagNoHeader {
		if err := printContextHeader(out, flagFull); err != nil {
			return err
		}
	}

	for _, ctx := range contexts {
		if err := printContext(out, &ctx, flagFull); err != nil {
			return err
		}
	}
	return nil
}
