package formatters

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
	"github.com/vapor-ware/synse-server-grpc/go"
)

const (
	// the pretty output format for `synse plugin capabilities` requests
	prettyCapabilities = "{{.Kind}}\t{{.Outputs}}\n"
)

// capabilitiesFormat collects the data that will be parsed into the output template.
type capabilitiesFormat struct {
	Kind    string
	Outputs string
}

// newCapabilitiesFormat is the handler for `synse plugin capabilities` commands that is
// used by the Formatter to add new capabilities data to the format context.
func newCapabilitiesFormat(data interface{}) (interface{}, error) {
	capability, ok := data.(*synse.DeviceCapability)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type *DeviceCapability", capability)
	}

	return &capabilitiesFormat{
		Kind:    capability.Kind,
		Outputs: fmt.Sprint(strings.Join(capability.Outputs, ", ")),
	}, nil
}

// NewCapabilitiesFormatter creates a new instance of a Formatter configured
// for the `synse plugin capabilities` command.
func NewCapabilitiesFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(c, newCapabilitiesFormat)
	f.Template = prettyCapabilities
	f.Decoder = &scheme.CapabilitiesOutput{}

	return f
}
