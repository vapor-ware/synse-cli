package formatters

import (
	"github.com/urfave/cli"
)

// NewStatusFormatter creates a new instance of a Formatter configured
// for a Synse Server status command.
func NewStatusFormatter(c *cli.Context, scheme interface{}) *Formatter {
	f := NewFormatter(
		c,
		&Formats{
			Yaml: scheme,
			JSON: scheme,
		},
	)
	return f
}
