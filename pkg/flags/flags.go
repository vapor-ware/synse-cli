package flags

import "github.com/urfave/cli"

// DebugFlag (--debug, -d) is the flag to enable debug logging output
var DebugFlag = cli.BoolFlag{
	Name:  "debug, d",
	Usage: "enable debug mode",
}

// ConfigFlag (--config) is the flag to display the YAML configuration for the CLI
var ConfigFlag = cli.BoolFlag{
	Name:  "config",
	Usage: "display the current CLI configuration",
}

// FormatFlag (--format) is the flag to specify the output format for a command.
// Note that not all commands support all formats.
var FormatFlag = cli.StringFlag{
	Name:  "format",
	Value: "pretty",
	Usage: "specify the output format for a command",
}

// NoHeaderFlag (--no-header) is a flag used to prevent printing the header
// when pretty printing. It has no effect when the output format is not 'pretty'.
var NoHeaderFlag = cli.BoolFlag{
	Name:  "no-header",
	Usage: "do not print the header when pretty-printing output",
}
