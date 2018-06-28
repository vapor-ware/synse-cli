package formatters

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
)

const (
	// the pretty output for plugins requests
	prettyPlugins = "{{.Tag}}\t{{.Network.Protocol}}\t{{.Network.Address}}\t{{.Version.Version}}\t{{.Health.Status}}\n"
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
			Name:        p.Name,
			Tag:         p.Tag,
			Description: p.Description,
			Maintainer:  p.Maintainer,
			VCS:         p.VCS,
			Network:     p.Network,
			Health:      p.Health,
			Version:     p.Version,
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
