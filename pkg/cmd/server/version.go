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
	"io"

	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

func init() {
	cmdVersion.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
	cmdVersion.Flags().BoolVarP(&flagJSON, "json", "", false, "print output as JSON")
	cmdVersion.Flags().BoolVarP(&flagYaml, "yaml", "", false, "print output as YAML")
}

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Display version information for the server",
	Long: utils.Doc(`
		Get version information for the configured Synse Server instance.

		The output of this command can be formatted as a table (default), as
		JSON, or as YAML. If specifying the output format, only one flag may
		be used. Using multiple output format flags will result in an error.

		For more information, see:
		<underscore>https://vapor-ware.github.io/synse-server/#version</>
	`),
	Run: func(cmd *cobra.Command, args []string) {
		// Error out if multiple output formats are specified.
		if flagJSON && flagYaml {
			exitutil.Err("cannot use multiple formatting flags at once")
		}

		exitutil.Err(serverVersion(cmd.OutOrStdout()))
	},
}

func serverVersion(out io.Writer) error {
	client, err := utils.NewSynseHTTPClient(flagContext, flagTLSCert)
	if err != nil {
		return err
	}

	response, err := client.Version()
	if err != nil {
		return err
	}

	printer := utils.NewPrinter(out, flagJSON, flagYaml, flagNoHeader)
	printer.SetHeader("VERSION", "API_VERSION")
	printer.SetRowFunc(serverVersionRowFunc)

	return printer.Write(response)
}
