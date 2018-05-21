package formatters

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
)

const (
	// the pretty output format for write requests
	prettyWrite = "{{.Transaction}}\t{{.Action}}\t{{$n := len .Raw}}{{range $i, $e := .Raw}}{{.}}{{if lt (plus1 $i) $n}}, {{end}}{{end}}\n"
)

// newWriteFormat is the handler for write commands that is used by the
// Formatter to add new write data to the format context.
func newWriteFormat(data interface{}) (interface{}, error) {
	write, ok := data.([]scheme.WriteTransaction)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type []scheme.WriteTransaction", data)
	}

	var out []interface{}
	for _, t := range write {
		out = append(out, &scheme.WriteOutput{
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
	f := NewFormatter(c, newWriteFormat)
	f.Template = prettyWrite
	f.Decoder = &scheme.WriteOutput{}

	return f
}
