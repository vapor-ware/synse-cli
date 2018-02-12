package server

import (
	"io"

	"fmt"

	"github.com/vapor-ware/synse-cli/formatters"
	"github.com/vapor-ware/synse-cli/scheme"
)

const (
	// the default output template for read requests
	readTmpl = "table {{.Reading}}\t{{.Value}}\n"
)

// readFormat collects the data that will be parsed into the output template.
type readFormat struct {
	Reading string
	Value   string
}

// newReadFormat is the handler for read commands that is used by the
// Formatter to add new read data to the format context.
func newReadFormat(data interface{}) (interface{}, error) {
	read, ok := data.(*scheme.Read)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type *scheme.Read", data)
	}

	var out []interface{}
	for readType, readData := range read.Data {
		out = append(out, &readFormat{
			Reading: readType,
			Value:   fmt.Sprintf("%v", readData.Value),
		})
	}

	return out, nil
}

// NewReadFormatter creates a new instance of a Formatter configured
// for the read command.
func NewReadFormatter(out io.Writer) *formatters.Formatter {
	f := formatters.NewFormatter(
		readTmpl,
		out,
	)
	f.SetHandler(newReadFormat)
	f.SetHeader(readFormat{
		Reading: "READING",
		Value:   "VALUE",
	})
	return f
}
