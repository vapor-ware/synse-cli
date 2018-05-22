package formatters

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
)

// NewStatusFormatter creates a new instance of a Formatter configured
// for a Synse Server status command.
func NewStatusFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(c, PassthroughHandler)
	f.Decoder = &scheme.TestStatus{}

	return f
}
