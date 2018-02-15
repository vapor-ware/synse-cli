package formatters

import (
	"fmt"
	"io"

	"github.com/vapor-ware/synse-cli/pkg/scheme"
	"github.com/vapor-ware/synse-cli/pkg/utils"
)

const (
	// the default output template for read requests
	readTmpl = "table {{.Type}}\t{{.Value}}\t{{.Timestamp}}\n"
)

// readFormat collects the data that will be parsed into the output template.
type readFormat struct {
	Type      string
	Value     string
	Timestamp string
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
			Type:      readType,
			Value:     fmt.Sprintf("%v", readData.Value),
			Timestamp: utils.ParseTimestamp(readData.Timestamp),
		})
	}

	return out, nil
}

// NewReadFormatter creates a new instance of a Formatter configured
// for the read command.
func NewReadFormatter(out io.Writer) *Formatter {
	f := NewFormatter(
		readTmpl,
		out,
	)
	f.SetHandler(newReadFormat)
	f.SetHeader(readFormat{
		Type:      "TYPE",
		Value:     "VALUE",
		Timestamp: "TIMESTAMP",
	})
	return f
}
