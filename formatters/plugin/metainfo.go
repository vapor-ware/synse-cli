package plugin

import (
	"fmt"
	"io"

	"github.com/vapor-ware/synse-cli/formatters"
	"github.com/vapor-ware/synse-server-grpc/go"
)

const (
	// the default output template for plugin metainfo requests
	metaTmpl = "table {{.ID}}\t{{.Type}}\t{{.Model}}\t{{.Protocol}}\t{{.Rack}}\t{{.Board}}\n"
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

// NewMetaFormatter creates a new instance of a Formatter configured
// for the plugin meta command.
func NewMetaFormatter(out io.Writer) *formatters.Formatter {
	f := formatters.NewFormatter(
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
