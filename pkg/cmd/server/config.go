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
	cmdConfig.Flags().BoolVarP(&flagYaml, "yaml", "", false, "print output as YAML")
}

var cmdConfig = &cobra.Command{
	Use:   "config",
	Short: "Display the configuration for the server",
	Long: utils.Doc(`
		Display the application configuration for the Synse Server instance.

		The output of this command can be formatted as JSON (default) or as
		YAML.

		For more information, see:
		<underscore>https://vapor-ware.github.io/synse-server/#config</>
	`),
	Run: func(cmd *cobra.Command, args []string) {
		exitutil.Err(serverConfig(cmd.OutOrStdout()))
	},
}

func serverConfig(out io.Writer) error {
	client, err := utils.NewSynseHTTPClient(flagContext, flagTLSCert)
	if err != nil {
		return err
	}

	response, err := client.Config()
	if err != nil {
		return err
	}

	printer := utils.NewPrinter(out, !flagYaml, flagYaml, flagNoHeader)
	return printer.Write(response)
}
