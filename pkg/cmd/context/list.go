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
	"io"
	"sort"

	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-cli/pkg/utils/exit"
	"github.com/vapor-ware/synse-cli/pkg/utils/sortable"
)

func init() {
	cmdList.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
	cmdList.Flags().BoolVarP(&flagJSON, "json", "", false, "print output as JSON")
	cmdList.Flags().BoolVarP(&flagYaml, "yaml", "", false, "print output as YAML")
}

var cmdList = &cobra.Command{
	Use:   "list",
	Short: "List all configured contexts",
	Long: utils.Doc(`
		List all configured contexts.

		This will display each configured context record.
	`),
	Aliases: []string{
		"ls",
	},
	Run: func(cmd *cobra.Command, args []string) {
		exiter := exit.FromCmd(cmd)

		// Error out if multiple output formats are specified.
		if flagJSON && flagYaml {
			exiter.Err("cannot use multiple formatting flags at once")
		}

		exiter.Err(listContexts(cmd.OutOrStdout()))
	},
}

func listContexts(out io.Writer) error {
	contexts := config.GetContexts()
	if len(contexts) == 0 {
		return nil
	}

	printer := utils.NewPrinter(out, flagJSON, flagYaml, flagNoHeader)
	printer.SetHeader("CURRENT", "NAME", "TYPE", "ADDRESS")
	printer.SetRowFunc(contextRowFunc)

	sort.Sort(sortable.ContextRecords(contexts))
	return printer.Write(contexts)
}
