package formatters

import (
	"github.com/urfave/cli"
)

// NewInfoFormatter creates a new instance of a Formatter configured
// for a Synse Server info command.
func NewInfoFormatter(c *cli.Context, scheme interface{}) *Formatter {
	f := NewFormatter(
		c,
		&Formats{
			Yaml: scheme,
			JSON: scheme,
		},
	)
	return f
}
