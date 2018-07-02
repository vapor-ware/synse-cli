package formatters

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
	"github.com/vapor-ware/synse-server-grpc/go"
)

const (
	// the pretty output format for `server capabilities` requests.
	prettyServerCapabilities = "{{.Plugin}}\t{{.Kind}}\t{{.Outputs}}\t\n"

	// the pretty output format for `plugin capabilities` requests.
	prettyPluginCapabilities = "{{.Kind}}\t{{.Outputs}}\t\n"
)

// serverCapabilityFormat is the format for `server capabilities` command.
type serverCapabilityFormat struct {
	Plugin  string
	Kind    string
	Outputs string
}

// pluginCapabilityFormat is the format for `plugin capabilities` command.
type pluginCapabilityFormat struct {
	Kind    string
	Outputs string
}

// newServerCapabilitiesFormat is the handler for `server capabilities` commands that is
// used by the Formatter to add new capabilities data to the format context.
func newServerCapabilitiesFormat(data interface{}) (interface{}, error) {
	capabilities, ok := data.([]scheme.Capability)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type []scheme.Capability", capabilities)
	}

	var out []interface{}
	for _, c := range capabilities {
		for _, d := range c.Devices {
			out = append(out, &serverCapabilityFormat{
				Plugin:  c.Plugin,
				Kind:    d.Kind,
				Outputs: fmt.Sprint(strings.Join(d.Outputs, ", ")),
			})
		}
	}
	return out, nil
}

// newPluginCapabilitiesFormat is the handler for `plugin capabilities` commands that is
// used by the Formatter to add new capabilities data to the format context.
func newPluginCapabilitiesFormat(data interface{}) (interface{}, error) {
	capability, ok := data.(*synse.DeviceCapability)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type *DeviceCapability", capability)
	}

	return &pluginCapabilityFormat{
		Kind:    capability.Kind,
		Outputs: fmt.Sprint(strings.Join(capability.Outputs, ", ")),
	}, nil
}

// NewServerCapabilitiesFormatter creates a new instance of a Formatter configured
// for the `server capabilities` command.
func NewServerCapabilitiesFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(c, newServerCapabilitiesFormat)
	f.Template = prettyServerCapabilities
	f.Decoder = &serverCapabilityFormat{}

	return f
}

// NewPluginCapabilitiesFormatter creates a new instance of a Formatter configured
// for the `plugin capabilities` command.
func NewPluginCapabilitiesFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(c, newPluginCapabilitiesFormat)
	f.Template = prettyPluginCapabilities
	f.Decoder = &pluginCapabilityFormat{}

	return f
}
