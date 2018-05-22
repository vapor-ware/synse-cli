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

	return &scheme.Transaction{
		ID:      transaction.ID,
		Context: transaction.Context,
		Status:  transaction.Status,
		State:   transaction.State,
		Created: utils.ParseTimestamp(transaction.Created),
		Updated: utils.ParseTimestamp(transaction.Updated),
		Message: transaction.Message,
	}, nil
}

// NewTransactionFormatter creates a new instance of a Formatter configured
// for the transaction command.
func NewTransactionFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(c, newTransactionFormat)
	f.Template = prettyTransaction
	f.Decoder = &scheme.Transaction{}

	return f
}
