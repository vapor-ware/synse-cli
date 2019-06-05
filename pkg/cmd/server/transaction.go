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
	"github.com/vapor-ware/synse-cli/pkg/utils/exit"
	"github.com/vapor-ware/synse-client-go/synse/scheme"
)

func init() {
	cmdTransaction.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
	cmdTransaction.Flags().BoolVarP(&flagJSON, "json", "", false, "print output as JSON")
	cmdTransaction.Flags().BoolVarP(&flagYaml, "yaml", "", false, "print output as YAML")
}

var cmdTransaction = &cobra.Command{
	Use:   "transaction [TRANSACTION...]",
	Short: "Check the status of transactions",
	Long: utils.Doc(`
		Check the status of write transactions.

		If no transaction(s) are specified by ID, all transactions are
		displayed.

		Writes in Synse are asynchronous. When a write is performed, a
		transaction is associated with the write and can be checked later
		to get the status of the write event. This command can be used to
		check that status.

		A transaction can have one of four statues:
		- <bold>PENDING</>: The write was received and is queued up, but has not
		    yet been executed.
		- <bold>WRITING</>: The write is in the process of being executed.
		- <bold>DONE</>: The write has completed successfully. This is a terminal
		    state. Once a transaction is in this state, it will no longer
		    be updated.
		- <bold>ERROR</>: An error has occurred at some point while trying to
		    execute the write. This is a terminal state. Once a transaction
		    is in this state, it will no longer be updated.

		The output of this command can be formatted as a table (default), as
		JSON, or as YAML. If specifying the output format, only one flag may
		be used. Using multiple output format flags will result in an error.

		For more information, see:
		<underscore>https://vapor-ware.github.io/synse-server/#transaction</>
	`),
	Run: func(cmd *cobra.Command, args []string) {
		exiter := exit.FromCmd(cmd)

		// Error out if multiple output formats are specified.
		if flagJSON && flagYaml {
			exiter.Err("cannot use multiple formatting flags at once")
		}

		exiter.Err(serverTransaction(cmd.OutOrStdout(), args))
	},
}

func serverTransaction(out io.Writer, transactions []string) error {
	client, err := utils.NewSynseHTTPClient(flagContext, flagTLSCert)
	if err != nil {
		return err
	}

	if len(transactions) == 0 {
		txns, err := client.Transactions()
		if err != nil {
			return err
		}
		for _, t := range txns {
			_, err = out.Write([]byte(t))
			if err != nil {
				return err
			}
		}
		return nil
	}

	var txns []*scheme.Transaction

	for _, t := range transactions {
		response, err := client.Transaction(t)
		if err != nil {
			return err
		}
		txns = append(txns, response)
	}

	if len(txns) == 0 {
		return nil
	}

	printer := utils.NewPrinter(out, flagJSON, flagYaml, flagNoHeader)
	printer.SetHeader("ID", "STATUS", "MESSAGE", "CREATED", "UPDATED")
	printer.SetRowFunc(serverTransactionRowFunc)

	return printer.Write(txns)
}
