package plugin

import (
	"fmt"
	"io"

	"github.com/vapor-ware/synse-cli/formatters"
	"github.com/vapor-ware/synse-server-grpc/go"
)

const (
	// the default output template for plugin reads
	readTmpl = "table {{.Type}}\t{{.Reading}}\t{{.Timestamp}}\n"
)

// readFormat collects the data that will be parsed into the output template.
type readFormat struct {
	Type      string
	Reading   string
	Timestamp string
}

// newReadFormat is the handler for plugin read commands that is used by the
// Formatter to add new read data to the format context.
func newReadFormat(data interface{}) (interface{}, error) {
	read, ok := data.(*synse.ReadResponse)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type *ReadResponse", read)
	}
	return &readFormat{
		Type:      read.Type,
		Reading:   read.Value,
		Timestamp: read.Timestamp,
	}, nil
}

// NewReadFormatter creates a new instance of a Formatter configured
// for the plugin read command.
func NewReadFormatter(out io.Writer) *formatters.Formatter {
	f := formatters.NewFormatter(
		readTmpl,
		out,
	)
	f.SetHandler(newReadFormat)
	f.SetHeader(readFormat{
		Type:      "TYPE",
		Reading:   "READING",
		Timestamp: "TIMESTAMP",
	})
	return f
}
