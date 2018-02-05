package server

import (
	"fmt"
	"net/http"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/client"
	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-cli/utils"
	"gopkg.in/yaml.v2"
)

// transactionURI
const transactionURI = "transaction"

// transactionCommand
var TransactionCommand = cli.Command{
	Name:     "transaction",
	Usage:    "transaction",
	Category: "Synse Server Actions",
	Action: func(c *cli.Context) error {
		return utils.CommandHandler(c, cmdTransaction(c))
	},
}

// cmdTransaction
func cmdTransaction(c *cli.Context) error {
	transactionID := c.Args().Get(0)
	if transactionID == "" {
		return cli.NewExitError("'transaction' requires 1 argument", 1)
	}

	transaction := &scheme.Transaction{}
	uri := fmt.Sprintf("%s/%s", transactionURI, transactionID)
	resp, err := client.New().Get(uri).ReceiveSuccess(transaction)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return err
	}

	out, err := yaml.Marshal(transaction)
	if err != nil {
		return err
	}
	fmt.Printf("%s", out)
	return nil
}
