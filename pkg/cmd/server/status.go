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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	"github.com/vapor-ware/synse-cli/pkg/utils/exit"
)

func init() {
	cmdStatus.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
	cmdStatus.Flags().BoolVarP(&flagJSON, "json", "", false, "print output as JSON")
	cmdStatus.Flags().BoolVarP(&flagYaml, "yaml", "", false, "print output as YAML")
}

var cmdStatus = &cobra.Command{
	Use:   "status",
	Short: "Get the status of the server instance",
	Long: utils.Doc(`
		Get the connectivity status for the configured Synse Server instance.

		The output of this command can be formatted as a table (default), as
		JSON, or as YAML. If specifying the output format, only one flag may
		be used. Using multiple output format flags will result in an error.

		For more information, see:
		<underscore>https://vapor-ware.github.io/synse-server/#test</>
	`),
	Run: func(cmd *cobra.Command, args []string) {
		exiter := exit.FromCmd(cmd)

		// Error out if multiple output formats are specified.
		if flagJSON && flagYaml {
			exiter.Err("cannot use multiple formatting flags at once")
		}

		exiter.Err(serverStatus(cmd.OutOrStdout()))
	},
}

func serverStatus(out io.Writer) error {
	log.Debug("creating new HTTP client")
	client, err := utils.NewSynseHTTPClient(flagContext, flagTLSCert)
	if err != nil {
		return err
	}

	log.Debug("issuing HTTP status request")
	response, err := client.Status()
	if err != nil {
		return err
	}

	printer := utils.NewPrinter(out, flagJSON, flagYaml, flagNoHeader)
	printer.SetHeader("STATUS", "TIMESTAMP")
	printer.SetRowFunc(serverStatusRowFunc)

	return printer.Write(response)
}
