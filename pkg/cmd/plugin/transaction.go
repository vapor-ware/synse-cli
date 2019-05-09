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

package plugin

import (
	"context"
	"encoding/json"
	"io"

	"github.com/spf13/cobra"
	"github.com/vapor-ware/synse-cli/pkg/utils"
	synse "github.com/vapor-ware/synse-server-grpc/go"
	"gopkg.in/yaml.v2"
)

func init() {
	cmdTransaction.Flags().BoolVarP(&flagNoHeader, "no-header", "n", false, "do not print out column headers")
	cmdTransaction.Flags().BoolVarP(&flagJson, "json", "", false, "print output as JSON")
	cmdTransaction.Flags().BoolVarP(&flagYaml, "yaml", "", false, "print output as YAML")
}

var cmdTransaction = &cobra.Command{
	Use:   "transaction [TRANSACTIONS...]",
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
	`),
	Run: func(cmd *cobra.Command, args []string) {
		// Error out if multiple output formats are specified.
		if flagJson && flagYaml {
			utils.Err("cannot use multiple formatting flags at once")
		}

		utils.Err(pluginTransaction(cmd.OutOrStdout(), args))
	},
}

// FIXME: this behemoth can be cleaned up & simplified..

func pluginTransaction(out io.Writer, transactions []string) error {
	conn, client, err := utils.NewSynseGrpcClient()
	if err != nil {
		return err
	}
	defer conn.Close()

	// -------- all transactions -------
	if len(transactions) == 0 {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		stream, err := client.Transactions(ctx, &synse.Empty{})
		if err != nil {
			return err
		}

		var txns []*synse.V3TransactionStatus
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
			txns = append(txns, resp)
		}

		if len(txns) == 0 {
			// TODO: on no reading, should it print a message "no readings",
			//   should it print nothing, or should it just print header info
			//   with no rows?
		}

		// Format output
		// FIXME: there is probably a way to clean this up / generalize this, but
		//   that can be done later.
		if flagJson {
			o, err := json.MarshalIndent(txns, "", "  ")
			if err != nil {
				return err
			}
			_, err = out.Write(append(o, '\n'))
			return err

		} else if flagYaml {
			o, err := yaml.Marshal(txns)
			if err != nil {
				return err
			}
			_, err = out.Write(o)
			return err

		} else {
			w := utils.NewTabWriter(out)
			defer w.Flush()

			if !flagNoHeader {
				if err := printTransactionStatusHeader(w); err != nil {
					return err
				}
			}

			for _, t := range txns {
				if err := printTransactionStatusRow(w, t); err != nil {
					return err
				}
			}
		}
		return nil
	}

	// ------ specified transactions --------
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var txns []*synse.V3TransactionStatus

	for _, transaction := range transactions {
		response, err := client.Transaction(ctx, &synse.V3TransactionSelector{
			Id: transaction,
		})
		if err != nil {
			return err
		}
		txns = append(txns, response)
	}

	if len(txns) == 0 {
		// TODO: on no reading, should it print a message "no readings",
		//   should it print nothing, or should it just print header info
		//   with no rows?
	}

	// Format output
	// FIXME: there is probably a way to clean this up / generalize this, but
	//   that can be done later.
	if flagJson {
		o, err := json.MarshalIndent(txns, "", "  ")
		if err != nil {
			return err
		}
		_, err = out.Write(append(o, '\n'))
		return err

	} else if flagYaml {
		o, err := yaml.Marshal(txns)
		if err != nil {
			return err
		}
		_, err = out.Write(o)
		return err

	} else {
		w := utils.NewTabWriter(out)
		defer w.Flush()

		if !flagNoHeader {
			if err := printTransactionStatusHeader(w); err != nil {
				return err
			}
		}

		for _, t := range txns {
			if err := printTransactionStatusRow(w, t); err != nil {
				return err
			}
		}
	}
	return nil
}
