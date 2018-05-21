package formatters

import (
	"github.com/urfave/cli"
)

// NewInfoFormatter creates a new instance of a Formatter configured
// for a Synse Server info command.
func NewInfoFormatter(c *cli.Context) *Formatter {
	f := NewFormatter(c, PassthroughHandler)

	return f
}
