package formatters

import (
	"github.com/urfave/cli"
	"github.com/vapor-ware/synse-cli/pkg/scheme"
)

// NewActiveFormatter creates a new instance of a Formatter configured
// for the host active command.
func NewActiveFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(c, PassthroughHandler)
	f.Decoder = &scheme.ActiveHostOutput{}

	return f
}
