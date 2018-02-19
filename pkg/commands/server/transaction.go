package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/client"
	"github.com/vapor-ware/synse-cli/pkg/formatters"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// transactionCmdName is the name for the 'transaction' command.
	transactionCmdName = "transaction"

	// transactionCmdUsage is the usage text for the 'transaction' command.
	transactionCmdUsage = "Check the state and status of a transaction"

	// transactionCmdDesc is the description for the 'transaction' command.
	transactionCmdDesc = `The transaction command hits the active Synse Server host's '/transaction'
	 endpoint, which returns the state and status of the specified transaction.`
)

// TransactionCommand is the CLI command for Synse Server's "transaction" API route.
var TransactionCommand = cli.Command{
	Name:        transactionCmdName,
	Usage:       transactionCmdUsage,
	Description: transactionCmdDesc,
	Category:    SynseActionsCategory,

	Action: func(c *cli.Context) error {
		return utils.CmdHandler(cmdTransaction(c))
	},
}

// cmdTransaction is the action for the TransactionCommand. It makes an "transaction"
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
