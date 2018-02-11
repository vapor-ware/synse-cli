package plugin

import (
	"fmt"
	"io"

	"github.com/vapor-ware/synse-cli/formatters"
	"github.com/vapor-ware/synse-server-grpc/go"
)

const (
	// the default output template for plugin transaction checks
	transactionTmpl = "table {{.Status}}\t{{.State}}\t{{.Created}}\t{{.Updated}}\n"
)

// transactionFormat collects the data that will be parsed into the output template.
type transactionFormat struct {
	Status  string
	State   string
	Created string
	Updated string
}

// newTransactionFormat is the handler for plugin transaction commands that is used by the
// Formatter to add new transaction data to the format context.
func newTransactionFormat(data interface{}) (interface{}, error) {
	transaction, ok := data.(*synse.WriteResponse)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type *WriteResponse", transaction)
	}

	return &transactionFormat{
		Status:  transaction.Status.String(),
		State:   transaction.State.String(),
		Created: transaction.Created,
		Updated: transaction.Updated,
	}, nil
}

// NewTransactionFormatter creates a new instance of a Formatter configured
// for the plugin transaction command.
func NewTransactionFormatter(out io.Writer) *formatters.Formatter {
	f := formatters.NewFormatter(
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
