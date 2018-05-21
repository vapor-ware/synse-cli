package formatters

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
)

// NewConfigFormatter creates a new instance of a Formatter configured
// for a Synse Server config command.
func NewConfigFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(c, PassthroughHandler)
	f.Decoder = &scheme.Config{}

	return f
}
