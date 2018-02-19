package formatters

import (
	"github.com/urfave/cli"
)

// NewActiveFormatter creates a new instance of a Formatter configured
// for the host active command.
func NewActiveFormatter(c *cli.Context, scheme interface{}) *Formatter {
	f := NewFormatter(
		c,
		&Formats{
			Yaml: scheme,
			JSON: scheme,
		},
	)
	return f
}
