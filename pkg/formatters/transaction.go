package formatters

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// the pretty output format for the transaction command
	prettyTransaction = "{{.Status}}\t{{.State}}\t{{.Created}}\t{{.Updated}}\t{{.Message}}\n"
)

// transactionFormat collects the data that will be parsed into the output template.
type transactionFormat struct {
	Status  string
	State   string
	Created string
	Updated string
	Message string
}

// newTransactionFormat is the handler for transaction commands that is used by the
// Formatter to add new transaction data to the format context.
func newTransactionFormat(data interface{}) (interface{}, error) {
	transaction, ok := data.(*scheme.Transaction)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type *scheme.Transaction", transaction)
	}

	// Set the message to "ok" if nothing is wrong
	if transaction.Message == "" {
		transaction.Message = "ok"
	}

	return &transactionFormat{
		Status:  transaction.Status,
		State:   transaction.State,
		Created: utils.ParseTimestamp(transaction.Created),
		Updated: utils.ParseTimestamp(transaction.Updated),
		Message: transaction.Message,
	}, nil
}

// NewTransactionFormatter creates a new instance of a Formatter configured
// for the transaction command.
func NewTransactionFormatter(c *cli.Context, data interface{}) *Formatter {
	f := NewFormatter(
		c,
		&Formats{
			Pretty: prettyTransaction,
			JSON: data,
			Yaml: data,
		},
	)
	f.SetHandler(newTransactionFormat)
	f.SetHeader(transactionFormat{
		Status:  "STATUS",
		State:   "STATE",
		Created: "CREATED",
		Updated: "UPDATED",
		Message: "MESSAGE",
	})
	return f
}
