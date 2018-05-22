package formatters

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
	"github.com/vapor-ware/synse-server-grpc/go"
)

const (
	// the pretty output format for scan requests
	prettyScan = "{{.Rack}}\t{{.Board}}\t{{.Device}}\t{{.Info}}\t{{.Type}}\n"

	// the pretty output format for plugin metainfo requests
	prettyMeta = "{{.ID}}\t{{.Type}}\t{{.Model}}\t{{.Protocol}}\t{{.Rack}}\t{{.Board}}\n"
)

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
	out, ok := data.([]interface{})
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type []*scheme.InternalScan", data)
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
func NewScanFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(c, newScanFormat)
	f.Template = prettyScan
	f.Decoder = &scheme.ScanDevice{}

	return f
}

// NewMetaFormatter creates a new instance of a Formatter configured
// for the plugin meta command.
func NewMetaFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(c, newMetaFormat)
	f.Template = prettyMeta
	f.Decoder = &scheme.MetaOutput{}

	return f
}
