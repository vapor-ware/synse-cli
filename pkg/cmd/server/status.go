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
	"encoding/json"
	"io"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

var cmdStatus = &cobra.Command{
	Use:   "status",
	Short: "Get the status of Synse Server",
	Long: heredoc.Doc(`
		Get the connectivity status of a Synse Server instance.

		This uses the '/test' endpoint, which is dependency-free and is used
		to determine whether the server instance is reachable or not. It does
		not provide any other indication of health.

		For more information on the status endpoint, see:
		https://vapor-ware.github.io/synse-server/#test
	`),

	Run: func(cmd *cobra.Command, args []string) {
		utils.Err(serverStatus(cmd.OutOrStdout()))
	},
}

func serverStatus(out io.Writer) error {
	client, err := utils.NewSynseHTTPClient()
	if err != nil {
		return err
	}

	response, err := client.Status()
	if err != nil {
		return err
	}

	// TODO: figure out output formatting
	o, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return err
	}
	_, err = out.Write(append(o, '\n'))
	return err
}
