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

package plugins

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

func init() {
	cmdInfo.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
	cmdInfo.Flags().BoolVarP(&flagJSON, "json", "", false, "print output as JSON")
	cmdInfo.Flags().BoolVarP(&flagYaml, "yaml", "", false, "print output as YAML")
}

var cmdInfo = &cobra.Command{
	Use:   "info PLUGIN",
	Short: "Display information about a registered plugin",
	Long: utils.Doc(`
		Display information about a plugin registered with the Synse Server
		instance.

		The output of this command can be formatted as a table (default), as
		JSON, or as YAML. If specifying the output format, only one flag may
		be used. Using multiple output format flags will result in an error.

		The default table view only provides a summary of the data. To see
		see the data in its entirety, use the JSON or YAML output formats.

		For more information, see:
		<underscore>https://vapor-ware.github.io/synse-server/#plugin</>
	`),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Error out if multiple output formats are specified.
		if flagJSON && flagYaml {
			exitutil.Err("cannot use multiple formatting flags at once")
		}

		exitutil.Err(serverPluginInfo(cmd.OutOrStdout(), args[0]))
	},
}

func serverPluginInfo(out io.Writer, plugin string) error {
	client, err := utils.NewSynseHTTPClient(flagContext, flagTLSCert)
	if err != nil {
		return err
	}

	response, err := client.Plugin(plugin)
	if err != nil {
		return err
	}

	printer := utils.NewPrinter(out, flagJSON, flagYaml, flagNoHeader)
	printer.SetHeader("ACTIVE", "ID", "TAG", "ADDRESS", "STATUS", "LAST_CHECK")
	printer.SetRowFunc(serverPluginRowFunc)

	return printer.Write(response)
}
