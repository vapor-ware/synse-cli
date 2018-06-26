package formatters

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
)

// const (
// 	// the pretty output format for write requests
// 	prettyWrite = "{{.Transaction}}\t{{.Context.Action}}\t{{$n := len .Context.Data}}{{range $i, $e := .Context.Data}}{{.}}{{if lt (plus1 $i) $n}}, {{end}}{{end}}\n"
// )

// newWriteFormat is the handler for write commands that is used by the
// Formatter to add new write data to the format context.
func newWriteFormat(data interface{}) (interface{}, error) {
	write, ok := data.([]scheme.WriteTransaction)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type []scheme.WriteTransaction", data)
	}

	var out []interface{}
	for _, t := range write {
		out = append(out, t)
	}
	return out, nil
}

// NewWriteFormatter creates a new instance of a Formatter configured
// for write command output.
func NewWriteFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(c, newWriteFormat)
	// f.Template = prettyWrite
	f.Decoder = &scheme.WriteTransaction{}

	return f
}
