package formatters

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
)

const (
	// the pretty output format for write requests
	prettyWrite = "{{.Transaction}}\t{{.Action}}\t{{range .Raw}}{{.}} {{end}}\n"
)

// writeFormat collects the data that will be parsed into the output template.
type writeFormat struct {
	Transaction string
	Action      string
	Raw         []string
}

// newWriteFormat is the handler for write commands that is used by the
// Formatter to add new write data to the format context.
func newWriteFormat(data interface{}) (interface{}, error) {
	write, ok := data.([]scheme.WriteTransaction)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type []scheme.WriteTransaction", data)
	}

	var out []interface{}
	for _, t := range write {
		out = append(out, &writeFormat{
			Transaction: t.Transaction,
			Action:      t.Context.Action,
			Raw:         t.Context.Raw,
		})
	}
	return out, nil
}

// NewWriteFormatter creates a new instance of a Formatter configured
// for write command output.
func NewWriteFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(
		c,
		&Formats{
			Pretty: prettyWrite,
		},
	)
	f.SetHandler(newWriteFormat)
	f.SetHeader(writeFormat{
		Transaction: "TRANSACTION ID",
		Action:      "ACTION",
		Raw:         []string{"RAW"},
	})
	return f
}
