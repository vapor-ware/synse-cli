package formatters

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
)

const (
	// the pretty output format for plugins requests
	prettyPlugins = "{{.Name}}\t{{.Network}}\t{{.Address}}\n"
)

// newPluginsFormat is the handler for plugins commands that is used by the
// Formatter to add new plugin data to the format context.
func newPluginsFormat(data interface{}) (interface{}, error) {
	plugins, ok := data.([]scheme.Plugin)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type []scheme.Plugin", data)
	}

	var out []interface{}
	for _, p := range plugins {
		out = append(out, &scheme.Plugin{
			Name:    p.Name,
			Network: p.Network,
			Address: p.Address,
		})
	}
	return out, nil
}

// NewPluginsFormatter creates a new instance of a Formatter configured
// for the plugins command.
func NewPluginsFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(c, newPluginsFormat)
	f.Template = prettyPlugins
	f.Decoder = &scheme.Plugin{}

	return f
}
