package formatters

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// the pretty output format for read requests
	prettyRead = "{{.Type}}\t{{.Value}}\t{{.Unit}}\t{{.Timestamp}}\n"
)

// newReadFormat is the handler for read commands that is used by the
// Formatter to add new read data to the format context.
func newReadFormat(data interface{}) (interface{}, error) {
	read, ok := data.(*scheme.Read)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type *scheme.Read", data)
	}

	var out []interface{}
	for readType, readData := range read.Data {
		out = append(out, &scheme.ReadOutput{
			Type:      readType,
			Value:     fmt.Sprintf("%v", readData.Value),
			Unit:      readData.Unit.Symbol,
			Timestamp: utils.ParseTimestamp(readData.Timestamp),
		})
	}

	return out, nil
}

// NewReadFormatter creates a new instance of a Formatter configured
// for the read command.
func NewReadFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(c, newReadFormat)
	f.Template = prettyRead
	f.Decoder = &scheme.ReadOutput{}

	return f
}
