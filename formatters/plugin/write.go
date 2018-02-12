package plugin

import (
	"fmt"
	"io"

	"github.com/vapor-ware/synse-cli/formatters"
	"github.com/vapor-ware/synse-server-grpc/go"
)

const (
	// the default output template for plugin writes
	writeTmpl = "table {{.ID}}\t{{.Action}}\t{{range .Raw}}{{.}} {{end}}\n"
)

// writeFormat collects the data that will be parsed into the output template.
type writeFormat struct {
	ID     string
	Action string
	Raw    []string
}

// newWriteFormat is the handler for plugin write commands that is used by the
// Formatter to add new write data to the format context.
func newWriteFormat(data interface{}) (interface{}, error) {
	write, ok := data.(*synse.Transactions)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type *Transactions", write)
	}

	var out []interface{}
	for id, ctx := range write.Transactions {
		var raw []string
		for _, r := range ctx.Raw {
			raw = append(raw, string(r))
		}

		out = append(out, &writeFormat{
			ID:     id,
			Action: ctx.Action,
			Raw:    raw,
		})
	}
	return out, nil
}

// NewWriteFormatter creates a new instance of a Formatter configured
// for the plugin write command.
func NewWriteFormatter(out io.Writer) *formatters.Formatter {
	f := formatters.NewFormatter(
		writeTmpl,
		out,
	)
	f.SetHandler(newWriteFormat)
	f.SetHeader(writeFormat{
		ID:     "TRANSACTION ID",
		Action: "ACTION",
		Raw:    []string{"RAW"},
	})
	return f
}
