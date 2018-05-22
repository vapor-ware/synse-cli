package formatters

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/config"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
)

const (
	// the pretty output format for host list command
	prettyList = "{{if .Active}}* {{else}}  {{end}}{{.Name}}\t{{.Address}}\n"
)

// newListFormat is the handler for host list commands that is used by the
// Formatter to add new list data to the format context.
func newListFormat(data interface{}) (interface{}, error) {
	cfg, ok := data.([]*config.HostConfig)
	if !ok {
		return nil, fmt.Errorf("formatter data %T not of type []*config.HostConfig", data)
	}
	var out []interface{}
	for _, c := range cfg {
		active := false
		if c.IsActiveHost() {
			active = true
		}
		out = append(out, &scheme.ListHostOutput{
			Active:  active,
			Name:    c.Name,
			Address: c.Address,
		})
	}
	return out, nil
}

// NewListFormatter creates a new instance of a Formatter configured
// for the host list command.
func NewListFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(c, newListFormat)
	f.Template = prettyList
	f.Decoder = &scheme.ListHostOutput{}

	return f
}
