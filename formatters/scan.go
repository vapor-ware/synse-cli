package formatters

import (
	"fmt"
	"io"

	"github.com/vapor-ware/synse-cli/scheme"
	"github.com/vapor-ware/synse-server-grpc/go"
)

const (
	// the default output template for scan requests
	scanTmpl = "table {{.Rack}}\t{{.Board}}\t{{.Device}}\t{{.Info}}\t{{.Type}}\n"

	// the default output template for plugin metainfo requests
	metaTmpl = "table {{.ID}}\t{{.Type}}\t{{.Model}}\t{{.Protocol}}\t{{.Rack}}\t{{.Board}}\n"
)

// scanFormat collects the data that will be parsed into the output template.
type scanFormat struct {
	Rack   string
	Board  string
	Device string
	Info   string
	Type   string
}

// metaFormat collects the data that will be parsed into the output template.
type metaFormat struct {
	ID       string
	Type     string
	Model    string
	Protocol string
	Rack     string
	Board    string
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

// newMetaFormat is the handler for plugin metainfo commands that is used by the
// Formatter to add new metainfo data to the format context.
func newMetaFormat(data interface{}) (interface{}, error) {
	meta, ok := data.(*synse.MetainfoResponse)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type *MetainfoResponse", meta)
	}
	return &metaFormat{
		ID:       meta.Uid,
		Type:     meta.Type,
		Model:    meta.Model,
		Protocol: meta.Protocol,
		Rack:     meta.Location.Rack,
		Board:    meta.Location.Board,
	}, nil
}

// NewScanFormatter creates a new instance of a Formatter configured
// for the scan command.
func NewScanFormatter(out io.Writer) *Formatter {
	f := NewFormatter(
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

// NewMetaFormatter creates a new instance of a Formatter configured
// for the plugin meta command.
func NewMetaFormatter(out io.Writer) *Formatter {
	f := NewFormatter(
		metaTmpl,
		out,
	)
	f.SetHandler(newMetaFormat)
	f.SetHeader(metaFormat{
		ID:       "ID",
		Type:     "TYPE",
		Model:    "MODEL",
		Protocol: "PROTOCOL",
		Rack:     "RACK",
		Board:    "BOARD",
	})
	return f
}
