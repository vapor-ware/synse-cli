package formatters

import (
	"fmt"
	"io"

	"github.com/vapor-ware/synse-cli/scheme"
)

const (
	// the default output template for the transaction command
	transactionTmpl = "table {{.Status}}\t{{.State}}\t{{.Created}}\t{{.Updated}}\n"
)

// transactionFormat collects the data that will be parsed into the output template.
type transactionFormat struct {
	Status  string
	State   string
	Created string
	Updated string
}

// newTransactionFormat is the handler for transaction commands that is used by the
// Formatter to add new transaction data to the format context.
func newTransactionFormat(data interface{}) (interface{}, error) {
	transaction, ok := data.(*scheme.Transaction)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type *scheme.Transaction", transaction)
	}

	return &transactionFormat{
		Status:  transaction.Status,
		State:   transaction.State,
		Created: transaction.Created,
		Updated: transaction.Updated,
	}, nil
}

// NewTransactionFormatter creates a new instance of a Formatter configured
// for the transaction command.
func NewTransactionFormatter(out io.Writer) *Formatter {
	f := NewFormatter(
		transactionTmpl,
		out,
	)
	f.SetHandler(newTransactionFormat)
	f.SetHeader(transactionFormat{
		Status:  "STATUS",
		State:   "STATE",
		Created: "CREATED",
		Updated: "UPDATED",
	})
	return f
}
