package formatters

import (
	"github.com/urfave/cli"
)

// NewConfigFormatter creates a new instance of a Formatter configured
// for a Synse Server config command.
func NewConfigFormatter(c *cli.Context, scheme interface{}) *Formatter {
	f := NewFormatter(
		c,
		&Formats{
			Yaml: scheme,
			JSON: scheme,
		},
	)
	return f
}
