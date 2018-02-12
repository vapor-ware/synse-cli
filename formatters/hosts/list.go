package hosts

import (
	"io"

	"fmt"

	"github.com/vapor-ware/synse-cli/config"
	"github.com/vapor-ware/synse-cli/formatters"
)

const (
	// the default output template for host list command
	listTmpl = "table {{if .Active}}* {{else}}  {{end}}{{.Name}}\t{{.Address}}\n"
)

type listFormat struct {
	Active  bool
	Name    string
	Address string
}

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
		out = append(out, &listFormat{
			Active:  active,
			Name:    c.Name,
			Address: c.Address,
		})
	}
	return out, nil
}

// NewListFormatter creates a new instance of a Formatter configured
// for the host list command.
func NewListFormatter(out io.Writer) *formatters.Formatter {
	f := formatters.NewFormatter(
		listTmpl,
		out,
	)
	f.SetHandler(newListFormat)
	return f
}
