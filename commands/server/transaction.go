package server

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/utils"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/client"
	"net/http"
	"fmt"
)

// transactionURI
const transactionURI = "transaction"

// transactionCommand
var transactionCommand = cli.Command{
	Name:    "transaction",
	Usage:   "transaction",
	Action: func(c *cli.Context) error {
		return utils.CommandHandler(c, cmdTransaction(c))
	},
}

// cmdTransaction
func cmdTransaction(c *cli.Context) error {
	transaction := &scheme.Transaction{}
	resp, err := client.New().Get(transactionURI).ReceiveSuccess(transaction)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return err
	}

	fmt.Println("unimplemented")
	return nil
}
