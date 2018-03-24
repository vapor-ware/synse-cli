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

// pluginsFormat collects the data that will be parsed into the output template.
type pluginsFormat struct {
	Name    string
	Network string
	Address string
}

// newPluginsFormat is the handler for plugins commands that is used by the
// Formatter to add new plugin data to the format context.
func newPluginsFormat(data interface{}) (interface{}, error) {
	plugins, ok := data.([]scheme.Plugin)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type []scheme.Plugin", data)
	}

	var out []interface{}
	for _, p := range plugins {
		out = append(out, &pluginsFormat{
			Name:    p.Name,
			Network: p.Network,
			Address: p.Address,
		})
	}
	return out, nil
}

// NewPluginsFormatter creates a new instance of a Formatter configured
// for the plugins command.
func NewPluginsFormatter(c *cli.Context, data interface{}) *Formatter {
	f := NewFormatter(
		c,
		&Formats{
			Pretty: prettyPlugins,
			JSON:   data,
			Yaml:   data,
		},
	)
	f.SetHandler(newPluginsFormat)
	f.SetHeader(pluginsFormat{
		Name:    "NAME",
		Network: "NETWORK",
		Address: "ADDRESS",
	})
	return f
}
