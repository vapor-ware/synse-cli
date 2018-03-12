package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/completion"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// transactionCmdName is the name for the 'transaction' command.
	transactionCmdName = "transaction"

	// transactionCmdUsage is the usage text for the 'transaction' command.
	transactionCmdUsage = "Check the state and status of a transaction"

	// transactionCmdArgsUsage is the argument usage for the 'transaction' command.
	transactionCmdArgsUsage = "TRANSACTION_ID"

	// transactionCmdDesc is the description for the 'transaction' command.
	transactionCmdDesc = `The transaction command hits the active Synse Server host's
  '/transaction' endpoint, which returns the state and status of
  the specified transaction.

  Writes in Synse Server are asynchronous, so to verify that a
  write has completed successfully, you must check the state of
  the transaction afterwards.

  The possible transaction states and statuses are:

  STATUS        STATE
  ----------    ----------
  unknown       ok
  pending       error
  writing
  done

Example:
  synse server transaction bah8volrogrg01o4sjtg

Formatting:
  The 'server transaction' command supports the following formatting
  options (via the CLI global --format flag):
    - pretty (default)`
)

// transactionCommand is the CLI command for Synse Server's "transaction" API route.
var transactionCommand = cli.Command{
	Name:        transactionCmdName,
	Usage:       transactionCmdUsage,
	Description: transactionCmdDesc,
	ArgsUsage:   transactionCmdArgsUsage,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdTransaction(c))
	},

	BashComplete: completion.CompleteTransactions,
}

// cmdTransaction is the action for the transactionCommand. It makes a "transaction"
// request against the active Synse Server instance.
func cmdTransaction(c *cli.Context) error {
	err := utils.RequiresArgsExact(1, c)
	if err != nil {
		return err
	}

	transactionID := c.Args().Get(0)

	transaction, err := client.Client.Transaction(transactionID)
	if err != nil {
		return err
	}

	formatter := formatters.NewTransactionFormatter(c)
	err = formatter.Add(transaction)
	if err != nil {
		return err
	}
	return formatter.Write()
}
