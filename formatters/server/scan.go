package server

import (
	"io"

	"fmt"

	"github.com/vapor-ware/synse-cli/formatters"
	"github.com/vapor-ware/synse-cli/scheme"
)

const (
	// the default output template for scan requests
	scanTmpl = "table {{.Rack}}\t{{.Board}}\t{{.Device}}\t{{.Info}}\t{{.Type}}\n"
)

// scanFormat collects the data that will be parsed into the output template.
type scanFormat struct {
	Rack   string
	Board  string
	Device string
	Info   string
	Type   string
}

// newScanFormat is the handler for scan commands that is used by the
// Formatter to add new scan data to the format context.
func newScanFormat(data interface{}) (interface{}, error) {
	scan, ok := data.([]*scheme.InternalScan)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type []*scheme.InternalScan", data)
	}

	var out []interface{}
	for _, item := range scan {
		out = append(out, item)
	}
	return out, nil
}

// NewScanFormatter creates a new instance of a Formatter configured
// for the scan command.
func NewScanFormatter(out io.Writer) *formatters.Formatter {
	f := formatters.NewFormatter(
		scanTmpl,
		out,
	)
	f.SetHandler(newScanFormat)
	f.SetHeader(scanFormat{
		Rack:   "RACK",
		Board:  "BOARD",
		Device: "DEVICE",
		Info:   "INFO",
		Type:   "TYPE",
	})
	return f
}
