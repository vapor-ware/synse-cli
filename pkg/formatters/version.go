package formatters

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
)

// NewVersionFormatter creates a new instance of a Formatter configured
// for a Synse Server version command.
func NewVersionFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(c, PassthroughHandler)
	f.Decoder = &scheme.Version{}

	return f
}
