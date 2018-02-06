package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/flags"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
)

// transactionBase is the base URI for the "transaction" route.
const transactionBase = "transaction"

// TransactionCommand is the CLI command for Synse Server's "transaction" API route.
var TransactionCommand = cli.Command{
	Name:     "transaction",
	Usage:    "Check the state and status of a transaction",
	Category: "Synse Server Actions",
	Flags: []cli.Flag{
		flags.OutputFlag,
	},
	Action: func(c *cli.Context) error {
		return utils.CmdHandler(c, cmdTransaction(c))
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

	transaction := &scheme.Transaction{}
	err = utils.DoGet(utils.MakeURI(transactionBase, transactionID), transaction)
	if err != nil {
		return err
	}

	return utils.FormatOutput(c, transaction)
}
