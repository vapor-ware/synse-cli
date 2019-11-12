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
	cmdInfo.Flags().BoolVarP(&flagYaml, "yaml", "", false, "print output as YAML")
}

var cmdInfo = &cobra.Command{
	Use:   "info DEVICE",
	Short: "Get details about a device",
	Long: utils.Doc(`
		Get details about a device.

		This command will get detailed information for the specified device,
		including its metadata, tags, read-write capabilities, and supported
		outputs.

		The output of this command can be formatted as JSON (default) or as
		YAML.

		For more information, see:
		<underscore>https://vapor-ware.github.io/synse-server/#info</>
	`),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		exit.FromCmd(cmd).Err(serverInfo(cmd.OutOrStdout(), args[0]))
	},
}

func serverInfo(out io.Writer, device string) error {
	log.Debug("creating new HTTP client")
	client, err := utils.NewSynseHTTPClient(flagContext, flagTLSCert)
	if err != nil {
		return err
	}

	log.WithField("device", device).Debug("issuing HTTP device info request")
	response, err := client.Info(device)
	if err != nil {
		return err
	}

	printer := utils.NewPrinter(out, !flagYaml, flagYaml, flagNoHeader)
	return printer.Write(response)
}
