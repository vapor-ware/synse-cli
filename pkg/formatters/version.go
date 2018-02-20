package formatters

import (
	"github.com/urfave/cli"
)

// NewVersionFormatter creates a new instance of a Formatter configured
// for a Synse Server version command.
func NewVersionFormatter(c *cli.Context, scheme interface{}) *Formatter {
	f := NewFormatter(
		c,
		&Formats{
			Yaml: scheme,
			JSON: scheme,
		},
	)
	return f
}
