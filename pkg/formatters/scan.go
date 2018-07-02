package formatters

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
	"github.com/vapor-ware/synse-server-grpc/go"
)

const (
	// the pretty output format for `synse server scan` requests
	prettyScan = "{{.Rack}}\t{{.Board}}\t{{.Device}}\t{{.Info}}\t{{.Type}}\n"

	// the pretty output format for `synse plugin devices` requests
	prettyDevices = "{{.ID}}\t{{.Kind}}\t{{.Plugin}}\t{{.Info}}\t{{.Rack}}\t{{.Board}}\n"
)

// devicesFormat collects the data that will be parsed into the output template.
type devicesFormat struct {
	ID     string
	Kind   string
	Plugin string
	Info   string
	Rack   string
	Board  string
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

// newDevicesFormat is the handler for plugin devices commands that is used by the
// Formatter to add new devices data to the format context.
func newDevicesFormat(data interface{}) (interface{}, error) {
	device, ok := data.(*synse.Device)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type *Device", device)
	}
	return &devicesFormat{
		ID:     device.Uid,
		Kind:   device.Kind,
		Plugin: device.Plugin,
		Info:   device.Info,
		Rack:   device.Location.Rack,
		Board:  device.Location.Board,
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

// NewDevicesFormatter creates a new instance of a Formatter configured
// for the plugin devices command.
func NewDevicesFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(c, newDevicesFormat)
	f.Template = prettyDevices
	f.Decoder = &scheme.DevicesOutput{}

	return f
}
