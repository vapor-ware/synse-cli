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

var cmdTransaction = &cobra.Command{
	Use:   "transaction",
	Short: "Check transaction state/status",
	Long: heredoc.Doc(`
		Check the state and status of a write transaction.

		Write in Synse Server are asynchronous. To verify the outcome
		of a write, this command can be used to check the transaction
		afterwards. Below are the possible states and statuses.

		States      Statuses
		 - ok        - unknown
		 - error     - pending
		             - writing
		             - done

		A transaction is considered complete either when it has reached
		the 'done' status or is in the error state.

		For more information, see:
		https://vapor-ware.github.io/synse-server/#transaction
	`),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils.Err(serverTransaction(cmd.OutOrStdout(), args[0]))
	},
}

func serverTransaction(out io.Writer, transaction string) error {
	client, err := utils.NewSynseHTTPClient()
	if err != nil {
		return err
	}

	response, err := client.Transaction(transaction)
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
