package formatters

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
)

const (
	// the pretty output for `server plugins` requests
	prettyServerPlugins = "{{.Tag}}\t{{.Protocol}}\t{{.Address}}\t{{.Version}}\t{{.Status}}\n"
)

// newServerPluginsFormat is the handler for `server plugins` commands that is used by the
// Formatter to add new plugin data to the format context.
func newServerPluginsFormat(data interface{}) (interface{}, error) {
	plugins, ok := data.([]scheme.Plugin)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type []scheme.Plugin", data)
	}

	var out []interface{}
	for _, p := range plugins {
		out = append(out, &scheme.ServerPluginOutput{
			Tag:      p.Tag,
			Protocol: p.Network.Protocol,
			Address:  p.Network.Address,
			Version:  p.Version.Version,
			Status:   p.Health.Status,
		})
	}
	return out, nil
}

// newServerPluginsInfoFormat is the handler for `server plugins info` commands that is used by the
// Formatter to add new plugin data to the format context.
func newServerPluginsInfoFormat(data interface{}) (interface{}, error) {
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

// NewServerPluginsFormatter creates a new instance of a Formatter configured
// for the `plugins` command.
func NewServerPluginsFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(c, newServerPluginsFormat)
	f.Template = prettyServerPlugins
	f.Decoder = &scheme.ServerPluginOutput{}

	return f
}

// NewServerPluginsInfoFormatter creates a new instance of a Formatter configured
// for the `server plugins info` command. The only difference between this function
// and the NewPluginsFormatter above is that, it doesn't use the pretty scheme
// to specify the returning field. It returns the metadata information of all
// available plugins instead.
func NewServerPluginsInfoFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(c, newServerPluginsInfoFormat)
	f.Decoder = &scheme.Plugin{}

	return f
}
