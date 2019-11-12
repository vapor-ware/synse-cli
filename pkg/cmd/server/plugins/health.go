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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-cli/pkg/utils/exit"
)

func init() {
	cmdHealth.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
	cmdHealth.Flags().BoolVarP(&flagJSON, "json", "", false, "print output as JSON")
	cmdHealth.Flags().BoolVarP(&flagYaml, "yaml", "", false, "print output as YAML")
}

var cmdHealth = &cobra.Command{
	Use:   "health",
	Short: "Display a summary of plugin health",
	Long: utils.Doc(`
		Display a summary of plugin health for the Synse Server instance.

		The output of this command can be formatted as a table (default), as
		JSON, or as YAML. If specifying the output format, only one flag may
		be used. Using multiple output format flags will result in an error.

		For more information, see:
		<underscore>https://vapor-ware.github.io/synse-server/#plugin-health</>
	`),
	Run: func(cmd *cobra.Command, args []string) {
		exiter := exit.FromCmd(cmd)

		// Error out if multiple output formats are specified.
		if flagJSON && flagYaml {
			exiter.Err("cannot use multiple formatting flags at once")
		}

		exiter.Err(serverPluginHealth(cmd.OutOrStdout()))
	},
}

func serverPluginHealth(out io.Writer) error {
	log.Debug("creating new HTTP client")
	client, err := utils.NewSynseHTTPClient(flagContext, flagTLSCert)
	if err != nil {
		return err
	}

	log.Debug("issuing HTTP plugin health request")
	response, err := client.PluginHealth()
	if err != nil {
		return err
	}

	printer := utils.NewPrinter(out, flagJSON, flagYaml, flagNoHeader)
	printer.SetHeader("STATUS", "HEALTHY", "UNHEALTHY", "ACTIVE", "INACTIVE")
	printer.SetRowFunc(serverPluginHealthRowFunc)

	return printer.Write(response)
}
